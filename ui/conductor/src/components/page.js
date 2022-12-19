import vuetify from '@/plugins/vuetify.js';
import {routeTitle} from '@/util/router.js';
import {computed} from 'vue';
import {useRoute} from 'vue-router/composables';

/**
 *
 * @return {*}
 */
export function usePage() {
  const currentRoute = useRoute();
  const themeColor = computed(() => {
    if (!currentRoute) return undefined;
    for (let i = currentRoute.matched.length - 1; i >= 0; i--) {
      const r = currentRoute.matched[i];
      const color = r.meta?.['logoColor'] || vuetify.framework.theme.currentTheme[r.name];
      if (color) return color;
    }
    // didn't find it
    return undefined;
  });
  const pageTitle = computed(() => {
    if (!currentRoute) return undefined;
    return routeTitle(currentRoute);
  });

  const hasSections = computed(() => currentRoute?.matched?.some(r => r.components?.sections));
  const hasNav = computed(() => currentRoute?.matched?.some(r => r.components?.nav));

  return {themeColor, pageTitle, hasSections, hasNav};
}
