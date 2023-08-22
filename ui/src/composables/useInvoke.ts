import { getInvocationResult } from "@/views/component/debug/service";
import { reactive, watch } from "vue";

export default function useInvokeResult(resContent: any) {
  const invokeResult = reactive<any>({ responseDefine: [], checkPoints: [] });
  const mockResult = {
    responseDefine: [
      {
        id: 0,
        responseCode: "",
        schema: "",
        codes: null,
        code: "",
        output: "",
        resultStatus: "pass",
        resultMsg: "随便写的错误信息",
        conditionId: 0,
        conditionEntityId: 0,
        conditionEntityType: "responseDefine",
        mediaType: "",
        component: "",
        invokeId: 0,
      },
    ],
    checkPoints: [
      {
        id: 0,
        type: "",
        expression: "",
        extractorVariable: "",
        operator: "",
        value: "",
        actualResult: "",
        resultStatus: "fail",
        resultMsg: "断言结果失败请求体错误了 a-jhdsjd",
        conditionId: 0,
        conditionEntityId: 0,
        conditionEntityType: "checkpoint",
        invokeId: 0,
      },
      {
        id: 0,
        type: "",
        expression: "",
        extractorVariable: "",
        operator: "",
        value: "",
        actualResult: "",
        resultStatus: "pass",
        resultMsg: "断言结果通过120-4",
        conditionId: 0,
        conditionEntityId: 0,
        conditionEntityType: "checkpoint",
        invokeId: 0,
      },
    ],
  };
  const getInvokeResult = async (invokeId: number) => {
    if (!invokeId) {
      return;
    }
    const res = await getInvocationResult(invokeId);
    if (res.code === 0) {
      Object.assign(invokeResult, {
        responseDefine: (res.data || []).filter(
          (e) => e.conditionEntityType === "conditionEntityType"
        ),
        checkPoints: (res.data || []).filter(
          (e) => e.conditionEntityType === "checkpoint"
        ),
      });
    }
  };

  watch(
    () => {
      return resContent.value;
    },
    (val) => {
      const { invokeId = 0 } = val || {};
      getInvokeResult(invokeId);
    },
    {
      immediate: true,
    }
  );

  return {
    invokeResult,
  };
}
