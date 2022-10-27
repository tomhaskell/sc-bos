package adapt

import (
	"context"
	"fmt"
	"github.com/vanti-dev/bsp-ew/internal/driver/bacnet/rpc"
	"github.com/vanti-dev/bsp-ew/internal/node"
	"github.com/vanti-dev/gobacnet"
	"github.com/vanti-dev/gobacnet/property"
	bactypes "github.com/vanti-dev/gobacnet/types"
	"github.com/vanti-dev/gobacnet/types/objecttype"
	"go.uber.org/multierr"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Device adapts a bacnet Device into a Smart Core traits and other apis.
func Device(name string, client *gobacnet.Client, device bactypes.Device) node.SelfAnnouncer {
	return &DeviceBacnetService{
		name:   name,
		client: client,
		device: device,
	}
}

// DeviceBacnetService implements rpc.BacnetDriverServiceServer targeting a single bacnet device.
// This provides read and write operations for objects on that device.
//
// We should keep this API up to date wrt the features available in gobacnet.Client.
type DeviceBacnetService struct {
	rpc.UnimplementedBacnetDriverServiceServer

	name   string
	client *gobacnet.Client
	device bactypes.Device
}

func (d *DeviceBacnetService) AnnounceSelf(a node.Announcer) node.Undo {
	return a.Announce(d.name, node.HasClient(rpc.WrapBacnetDriverService(d)))
}

func (d *DeviceBacnetService) ReadProperty(ctx context.Context, request *rpc.ReadPropertyRequest) (*rpc.ReadPropertyResponse, error) {
	if request.ObjectIdentifier == nil {
		return nil, status.Errorf(codes.InvalidArgument, "missing object_identifier")
	}

	readProperty, err := d.client.ReadProperty(d.device, bactypes.ReadPropertyData{
		Object: bactypes.Object{
			ID:         d.objectIDFromProto(request.ObjectIdentifier),
			Properties: []bactypes.Property{d.propertyFromProtoForRead(request.PropertyReference)},
		},
	})
	if err != nil {
		return nil, err
	}

	if readProperty.ErrorCode != 0 {
		return nil, status.Errorf(codes.Internal, "Error(%d) from BACnet device", readProperty.ErrorCode)
	}

	result, err := d.propertyToProtoReadResult(readProperty.Object.Properties[0])
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Reading property %v", err)
	}
	return &rpc.ReadPropertyResponse{
		ObjectIdentifier: d.objectIDToProto(readProperty.Object.ID),
		Result:           result,
	}, nil
}

func (d *DeviceBacnetService) ReadPropertyMultiple(ctx context.Context, request *rpc.ReadPropertyMultipleRequest) (*rpc.ReadPropertyMultipleResponse, error) {
	bacReq := bactypes.ReadMultipleProperty{}
	for _, spec := range request.ReadSpecifications {
		obj := bactypes.Object{
			ID: d.objectIDFromProto(spec.ObjectIdentifier),
		}
		if len(spec.PropertyReferences) == 0 {
			obj.Properties = append(obj.Properties, d.defaultReadProperty())
		} else {
			for _, prop := range spec.PropertyReferences {
				obj.Properties = append(obj.Properties, d.propertyFromProtoForRead(prop))
			}
		}
		bacReq.Objects = append(bacReq.Objects, obj)
	}
	readProperties, err := d.client.ReadMultiProperty(d.device, bacReq)
	if err != nil {
		return nil, err
	}
	if readProperties.ErrorCode != 0 {
		return nil, status.Errorf(codes.Internal, "Error(%d) from BACnet device", readProperties.ErrorCode)
	}

	res := &rpc.ReadPropertyMultipleResponse{}
	for _, object := range readProperties.Objects {
		item := &rpc.ReadPropertyMultipleResponse_ReadResult{
			ObjectIdentifier: d.objectIDToProto(object.ID),
		}
		for _, p := range object.Properties {
			result, e := d.propertyToProtoReadResult(p)
			if e != nil {
				err = multierr.Append(err, e)
				continue
			}
			item.Results = append(item.Results, result)
		}
		res.ReadResults = append(res.ReadResults, item)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Reading properties %v", err)
	}
	return res, nil
}

func (d *DeviceBacnetService) WriteProperty(ctx context.Context, request *rpc.WritePropertyRequest) (*rpc.WritePropertyResponse, error) {
	writeProp, err := d.propertyFromProtoForWrite(request.WriteValue)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Interpreting request %v", err)
	}
	data := bactypes.ReadPropertyData{
		Object: bactypes.Object{
			ID:         d.objectIDFromProto(request.ObjectIdentifier),
			Properties: []bactypes.Property{writeProp},
		},
	}
	return &rpc.WritePropertyResponse{}, d.client.WriteProperty(d.device, data, uint(request.WriteValue.Priority))
}

func (d *DeviceBacnetService) WritePropertyMultiple(ctx context.Context, request *rpc.WritePropertyMultipleRequest) (*rpc.WritePropertyMultipleResponse, error) {
	// client doesn't implement WritePropertyMultiple! :(
	return d.UnimplementedBacnetDriverServiceServer.WritePropertyMultiple(ctx, request)
}

func (d *DeviceBacnetService) objectIDFromProto(identifier *rpc.ObjectIdentifier) bactypes.ObjectID {
	return bactypes.ObjectID{
		Type:     objecttype.ObjectType(identifier.Type),
		Instance: bactypes.ObjectInstance(identifier.Instance),
	}
}

func (d *DeviceBacnetService) objectIDToProto(id bactypes.ObjectID) *rpc.ObjectIdentifier {
	return &rpc.ObjectIdentifier{
		Type:     uint32(id.Type),
		Instance: uint32(id.Instance),
	}
}

func (d *DeviceBacnetService) propertyFromProtoForRead(reference *rpc.PropertyReference) bactypes.Property {
	if reference == nil {
		return d.defaultReadProperty()
	}
	p := bactypes.Property{
		ID:         property.ID(reference.Identifier),
		ArrayIndex: bactypes.ArrayAll,
	}
	if reference.ArrayIndex != nil {
		p.ArrayIndex = *reference.ArrayIndex
	}
	return p
}

func (d *DeviceBacnetService) defaultReadProperty() bactypes.Property {
	return bactypes.Property{
		ID:         property.PresentValue,
		ArrayIndex: bactypes.ArrayAll,
	}
}

func (d *DeviceBacnetService) propertyFromProtoForWrite(reference *rpc.PropertyWriteValue) (bactypes.Property, error) {
	p := d.propertyFromProtoForRead(reference.PropertyReference)
	var err error
	p.Data, err = d.propertyValueFromProto(reference.Value)
	return p, err
}

func (d *DeviceBacnetService) propertyToProto(p bactypes.Property) *rpc.PropertyReference {
	res := &rpc.PropertyReference{
		Identifier: uint32(p.ID),
	}
	if p.ArrayIndex != bactypes.ArrayAll {
		res.ArrayIndex = &p.ArrayIndex
	}
	return res
}

func (d *DeviceBacnetService) propertyValueToProto(p bactypes.Property) (*rpc.PropertyValue, error) {
	// Property.Data doesn't distinguish between 8, 16, 32 bit data, they are all 32.
	// It also doesn't support 64 bit data, so we don't either.
	// It also doesn't support bit_string, it returns an error that we should catch earlier in the request.
	// See gobacnet/encoding/appdata.go:240 - Decoder.AppData - for the supported types

	switch v := p.Data.(type) {
	case bactypes.Null:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Null{Null: true}}, nil
	case bool:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Boolean{Boolean: v}}, nil
	case uint32:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Unsigned32{Unsigned32: v}}, nil
	case int32:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Integer32{Integer32: v}}, nil
	case float32:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Real{Real: v}}, nil
	case float64:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Double{Double: v}}, nil
	case []byte:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_OctetString{OctetString: v}}, nil
	case string:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_CharacterString{CharacterString: v}}, nil
	case bactypes.Date:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Date{Date: d.dateToProto(v)}}, nil
	case bactypes.Time:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_Time{Time: d.timeToProto(v)}}, nil
	case bactypes.ObjectID:
		return &rpc.PropertyValue{Value: &rpc.PropertyValue_ObjectIdentifier{ObjectIdentifier: d.objectIDToProto(v)}}, nil
	}

	return nil, fmt.Errorf("unknown bacnet type %v (%T)", p.Data, p.Data)
}

func (d *DeviceBacnetService) propertyValueFromProto(p *rpc.PropertyValue) (any, error) {
	switch v := p.Value.(type) {
	case *rpc.PropertyValue_Null:
		return bactypes.Null{}, nil
	case *rpc.PropertyValue_Boolean:
		return v.Boolean, nil
	case *rpc.PropertyValue_Unsigned32:
		return v.Unsigned32, nil
	case *rpc.PropertyValue_Unsigned64:
		return v.Unsigned64, nil
	case *rpc.PropertyValue_Integer32:
		return v.Integer32, nil
	case *rpc.PropertyValue_Integer64:
		return v.Integer64, nil
	case *rpc.PropertyValue_Real:
		return v.Real, nil
	case *rpc.PropertyValue_Double:
		return v.Double, nil
	case *rpc.PropertyValue_OctetString:
		return v.OctetString, nil
	case *rpc.PropertyValue_CharacterString:
		return v.CharacterString, nil
	case *rpc.PropertyValue_BitString:
		// not supported
	case *rpc.PropertyValue_Enumerated:
		return bactypes.Enumerated(v.Enumerated), nil
	case *rpc.PropertyValue_Date:
		return d.dateFromProto(v.Date), nil
	case *rpc.PropertyValue_Time:
		return d.timeFromProto(v.Time), nil
	case *rpc.PropertyValue_ObjectIdentifier:
		return d.objectIDFromProto(v.ObjectIdentifier), nil
	}

	return nil, fmt.Errorf("unknown proto oneof type %v (%T)", p.Value, p.Value)
}

func (d *DeviceBacnetService) dateToProto(date bactypes.Date) *rpc.PropertyValue_DateValue {
	return &rpc.PropertyValue_DateValue{
		Year:       uint32(date.Year),
		Month:      uint32(date.Month),
		DayOfMonth: uint32(date.Day),
		DayOfWeek:  uint32(date.DayOfWeek),
	}
}

func (d *DeviceBacnetService) dateFromProto(date *rpc.PropertyValue_DateValue) bactypes.Date {
	return bactypes.Date{
		Year:      int(date.Year),
		Month:     int(date.Month),
		Day:       int(date.DayOfMonth),
		DayOfWeek: bactypes.DayOfWeek(date.DayOfWeek),
	}
}

func (d *DeviceBacnetService) timeToProto(time bactypes.Time) *rpc.PropertyValue_TimeValue {
	octetToProto := func(o int) *uint32 {
		if o == bactypes.UnspecifiedTime {
			return nil
		}
		ui := uint32(o)
		return &ui
	}
	return &rpc.PropertyValue_TimeValue{
		Hour:               octetToProto(time.Hour),
		Minute:             octetToProto(time.Minute),
		Second:             octetToProto(time.Second),
		HundredthsOfSecond: octetToProto(time.Millisecond / 10),
	}
}

func (d *DeviceBacnetService) timeFromProto(time *rpc.PropertyValue_TimeValue) bactypes.Time {
	octetFromProto := func(p *uint32) int {
		if p == nil {
			return bactypes.UnspecifiedTime
		}
		return int(*p)
	}
	t := bactypes.Time{
		Hour:        octetFromProto(time.Hour),
		Minute:      octetFromProto(time.Minute),
		Second:      octetFromProto(time.Second),
		Millisecond: octetFromProto(time.HundredthsOfSecond),
	}
	if t.Millisecond != bactypes.UnspecifiedTime {
		t.Millisecond *= 10
	}
	return t
}

func (d *DeviceBacnetService) propertyToProtoReadResult(p bactypes.Property) (*rpc.PropertyReadResult, error) {
	value, err := d.propertyValueToProto(p)
	if err != nil {
		return nil, err
	}
	return &rpc.PropertyReadResult{
		PropertyReference: d.propertyToProto(p),
		Value:             value,
	}, nil
}