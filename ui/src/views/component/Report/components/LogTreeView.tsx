import {defineComponent, ref, watch,} from 'vue';
import './LogTreeView.less';
import ScenarioHeader from "./ScenarioHeader.vue";
import InterfaceHeader from "./InterfaceHeader.vue";
import ProcessorHeader from "./ProcessorHeader.vue";
import InterfaceContent from "./InterfaceContent.vue";
import LogContent from "./LogContent.vue";
import {showArrowScenarioType,} from "@/views/scenario/components/Design/config";

export default defineComponent({
    name: 'LogTreeView',
    props: {
        treeData: Array,
        isSingleScenario: Boolean,
    },
    emits: ['change'],
    setup(props, {emit}) {
        const activeKeyMap = ref({});
        // 加载数据时，迭代器每 10 页，分一页
        const pageInfo = ref({});

        function change(uid, keys) {
            activeKeyMap.value[uid] = keys;
        }

        function changePageInfo(pid) {
            pageInfo.value[pid] = (pageInfo.value[pid] || 10) + 10;
        }

        /**
         * 点击更多
         * */
        function handleMore() {
            console.log('more')
        }

        watch(() => props.treeData, (newVal: any) => {
            if (newVal?.length) {
                newVal.forEach((item) => {
                    activeKeyMap.value[item.id] = [item.id];
                })
            }
        }, {immediate: true, deep: true})

        /**
         * @desc 渲染场景执行树
         * @param logs 需要渲染的场景类型
         * @param source 源数据
         * */
        function renderScenario(logs: any, source: any) {
            if (!logs) return null;

            function renderHeader(log) {
                // 接口场景
                if (log.processorCategory === 'processor_interface') {
                    return <InterfaceHeader endpointData={log}/>
                }
                // 其他场景
                return <ProcessorHeader record={log} onMore={handleMore}/>
            }

            function renderContent(log) {
                if (log.processorCategory === 'processor_interface' && log.detail !== undefined && log.detail !== '{}') {
                    return <InterfaceContent endpointData={log}/>
                }
                // 场景中的叶子节点不再渲染
                if (!showArrowScenarioType.includes(log.processorType)) {
                    return null;
                }
                return null;
            }

            // 渲染场景的标题，主要针对迭代场景，需要标识当前迭代的序号
            function renderCollapseTitle(log, logIndex, srcLog) {
                if (!log?.round) {
                    return null;
                }
                return (<span class={'collapse-title'}>{log?.round}</span>)
            }

            // 渲染分页
            function renderPage(pid) {
                return (<div class={'more-btn'}>
                    <a-button onClick={() => {
                        changePageInfo(pid);
                    }} type="link">加载更多</a-button>
                </div>)
            }

            const renderLogs = (log) => {
                if (!log?.id) {
                    return;
                }
                return <a-collapse-panel
                    key={log.id}
                    header={renderHeader(log)}
                    showArrow={showArrowScenarioType.includes(log.processorType) || (log.processorCategory === 'processor_interface' && log.detail !== undefined && log.detail !== '{}')}>
                    {renderContent(log)}
                    {
                        log?.logs?.map((item, itemIndex, srcLog) => {
                            const pid = item.parentLogId || item.parentId;
                            const pageSize = pageInfo.value[pid] || 10;
                            const isHideLog = itemIndex > pageSize;
                            if (isHideLog) {
                                return null
                            }
                            return <div
                                class={[item.processorType === 'processor_logic_else' ? 'log-item-else' : 'log-item', itemIndex === 0 ? 'log-item-first' : '']}>
                                {renderCollapseTitle(item, itemIndex, log)}
                                <a-collapse>
                                    {renderLogs(item)}
                                </a-collapse>
                                {pageSize === itemIndex + 1 ? renderPage(pid) : null}
                            </div>
                        })
                    }
                </a-collapse-panel>;
            };
            return logs.map((log, logIndex) => {
                return <div key={log.id}
                            class={[log.processorType === 'processor_logic_else' ? 'log-item-else' : 'log-item', logIndex === 0 ? 'log-item-first' : '']}>
                    <a-collapse>
                        {renderLogs(log)}
                    </a-collapse>
                </div>
            })
        }

        // 渲染场景，一级目录, 即场景列表
        function renderScenarioList(list) {
            console.log('renderScenarioList +++', list)
            if (!list?.length) {
                return null
            }
            // 如果是单场景，直接渲染场景
            if (list.length === 1 && props.isSingleScenario) {
                return renderScenario(list[0]?.logs, list[0])
            }

            const renderHeader = (item) => {
                return <ScenarioHeader record={item}/>
            }

            return list.map((item, index) => {
                return <div class={'scenario-item'} key={item.id}>
                    <a-collapse
                        activeKey={activeKeyMap.value[item.id]}
                        onChange={(key) => {
                            change(item.id, key)
                        }}>
                        <a-collapse-panel key={item.id} header={renderHeader(item)}>
                            {renderScenario(item?.logs, item)}
                        </a-collapse-panel>
                    </a-collapse>
                </div>
            })
        }

        return () => (
            <div class={'log-tree-view dp-log-tree-view'}>
                {renderScenarioList(props.treeData)}
            </div>
        )
    }
})
