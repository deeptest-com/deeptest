<template>
  <div class="flex flex-col flex-1">
    <SmartTabs styles="sticky bg-primary top-upperPrimaryStickyFold z-10">
      <SmartTab
          :id="'params'"
          :label="`${$t('tab.parameters')}`"
          :selected="true"
          :info="`${newActiveParamsCount$}`"
      >
        <HttpParameters />
      </SmartTab>

      <SmartTab :id="'bodyParams'" :label="`${$t('tab.body')}`">
        <HttpBody />
      </SmartTab>

      <SmartTab
          :id="'headers'"
          :label="`${$t('tab.headers')}`"
          :info="`${newActiveHeadersCount$}`"
      >
        <HttpHeaders />
      </SmartTab>

      <SmartTab
          :id="'authorization'"
          :label="`${$t('tab.authorization')}`"
      >
        <HttpAuthorization />
      </SmartTab>

      <SmartTab
          :id="'preRequestScript'"
          :label="`${$t('tab.pre_request_script')}`"
          :indicator="
                preRequestScript && preRequestScript.length > 0 ? true : false
              "
      >
        <HttpPreRequestScript />
      </SmartTab>

      <SmartTab
          :id="'tests'"
          :label="`${$t('tab.tests')}`"
          :indicator="testScript && testScript.length > 0 ? true : false"
      >
        <HttpTests />
      </SmartTab>
    </SmartTabs>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent } from "@nuxtjs/composition-api"
import { useReadonlyStream } from "~/helpers/utils/composables"
import { restResponse$ } from "~/newstore/RESTSession"

export default defineComponent({
  setup() {
    const response = useReadonlyStream(restResponse$, null)

    const hasResponse = computed(
      () =>
        response.value?.type === "success" || response.value?.type === "fail"
    )

    const loading = computed(
      () => response.value === null || response.value.type === "loading"
    )

    return {
      hasResponse,
      response,
      loading,
    }
  },
})
</script>
