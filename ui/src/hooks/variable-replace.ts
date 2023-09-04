import { ref, onMounted, onUnmounted, Ref } from 'vue'
import {getContextMenuStyle2} from "@/utils/dom";
import bus from "@/utils/eventBus";
import settings from "@/config/settings";

interface VariableReplacement {
    showContextMenu: Ref<boolean>,
    contextMenuStyle: Ref,

    onContextMenuShow: Function,
    onMenuClick: Function,
}

function useVariableReplace(src): VariableReplacement {
    const showContextMenu = ref(false)
    const paramIndex = ref(-1)
    const contextTarget = ref({} as any)
    const contextMenuStyle = ref({} as any)

    const onContextMenuShow = (idx, e) => {
        console.log('onContextMenuShow in hook', idx, e.target)
        if (!e) return

        contextMenuStyle.value = getContextMenuStyle2(e)
        contextTarget.value = e.target
        paramIndex.value = idx

        showContextMenu.value = true
    }
    const onMenuClick = (key) => {
        console.log('onMenuClick in hook', key)

        if (key === 'use-variable') {
            bus.emit(settings.eventVariableSelectionStatus, {
                src: src,
                index: paramIndex.value,
                data: contextTarget,
            });
        }

        showContextMenu.value = false
    }

    onMounted(() => {
        console.log('useVariableReplace onMounted')
    })

    onUnmounted(() => {
        console.log('useVariableReplace onUnmounted')
    })

    return {
        showContextMenu,
        contextMenuStyle,
        onContextMenuShow,
        onMenuClick
    } as VariableReplacement
}

export default useVariableReplace