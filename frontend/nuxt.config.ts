// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  future: {
    compatibilityVersion: 4,
  },
  devtools: { enabled: true },
  typescript: {
    typeCheck: true
  },
  modules: ['@nuxtjs/tailwindcss', '@nuxt/icon'],
  experimental: {
      scanPageMeta: 'after-resolve',
      sharedPrerenderData: false,
      compileTemplate: true,
      resetAsyncDataToUndefined: true,
      templateUtils: true,
      relativeWatchPaths: true,
      normalizeComponentNames: false,
      spaLoadingTemplateLocation: 'within',
      parseErrorData: false,
      pendingWhenIdle: true,
      alwaysRunFetchOnKeyChange: true,
      defaults: {
        useAsyncData: {
          deep: true
        }
      }
    },
    features: {
      inlineStyles: true
    },
    unhead: {
      renderSSRHeadOptions: {
        omitLineBreaks: false
      }
    }
});