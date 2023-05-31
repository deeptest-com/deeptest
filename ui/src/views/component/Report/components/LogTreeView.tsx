import {defineComponent, ref, watch,} from 'vue';
import './LogTreeView.less';
import {DownOutlined, PlusOutlined, RightOutlined,} from '@ant-design/icons-vue';
import Actions from "./Actions.vue";
import ScenarioHeader from "./ScenarioHeader.vue";

import EndpointHeader from "./EndpointHeader.vue";
import EndpointContent from "./EndpointContent.vue";

export default defineComponent({
    name: 'LogTreeView',
    props: {
        treeData: Array,
    },
    emits: ['change'],
    setup(props, {emit}) {
        const activeKey = ref([410, 411, 41, 1351, 1353,]);
        function changeActivekey(keys) {
            console.log('832', keys)
        }
        /**
         * @desc 渲染场景执行树
         * @param logs 需要渲染的场景类型
         * @param source 源数据
         * */
        function renderScenario(logs: any, source: any) {
            if (!logs) return null;

            function renderHeader(log) {
                if(log.processorCategory === 'processor_interface'){
                    return <EndpointHeader endpointData={log}/>
                }
                return <span>{log.name}：{log.summary}</span>
            }

            function renderContent(log) {
                if(log.processorCategory === 'processor_interface'){
                    return <EndpointContent endpointData={log}/>
                }
                if(log.processorCategory === 'processor_action'){
                    return null;
                }
                return null;
            }
            const renderLogs = (log) => {
                if (!log?.id) {
                    return;
                }
                return <a-collapse-panel header={renderHeader(log)}>
                    {renderContent(log)}
                    {
                        log?.logs?.map((log) => {
                            return <div class={'log-item'}>
                                <a-collapse>
                                    {renderLogs(log)}
                                </a-collapse>
                            </div>
                        })
                    }
                </a-collapse-panel>;
            };
            return logs.map((log) => {
                return <div class={'log-item'}>
                    <a-collapse>
                        {renderLogs(log)}
                    </a-collapse>
                </div>
            })
        }

        watch(() => props.treeData, (newVal) => {
            console.log('333333 newVal', newVal)
        })

        // 渲染场景，一级目录
        function renderScenarioList(list) {
            if (!list?.length) {
                return null
            }

            const renderHeader = (item) => {
                return <ScenarioHeader record={item}/>
            }

            return list.map((item, index) => {
                console.log(item.name)
                return <div class={'scenario-item'}>
                    <a-collapse>
                        <a-collapse-panel header={renderHeader(item)}>
                            {renderScenario(item?.logs?.[0]?.logs, item)}
                        </a-collapse-panel>
                    </a-collapse>
                </div>
            })
        }

        return () => (
            <div class={'log-tree-view'}>
                {renderScenarioList(props.treeData)}
            </div>
        )
    }
})
