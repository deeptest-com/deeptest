import {defineComponent, ref, watch,} from 'vue';
import './schema.less';
import {DownOutlined, PlusOutlined, RightOutlined,} from '@ant-design/icons-vue';
import SplitDivider from "./SplitDivider.vue";
import ExtraInfo from "./ExtraInfo.vue";
import DataTypeSetting from './DataTypeSetting.vue';
import cloneDeep from "lodash/cloneDeep";
import {
    addExtraViewInfo,
    findLastNotArrayNode,
    generateSchemaByArray,
    isArray,
    isNormalType,
    isObject,
    isRef,
    removeExtraViewInfo,
} from './utils';
import {
    treeLevelWidth
} from './config';
import {message} from "ant-design-vue";
import {useStore} from "vuex";
import {StateType as ServeStateType} from "@/store/serve";

export default defineComponent({
    name: 'SchemeEditor',
    props: {
        value: Object,
        contentStyle: Object,
        serveId: Number,
        refsOptions: Array,
        components: Array,
    },
    emits: ['change'],
    setup(props, {emit}) {
        const store = useStore<{ Endpoint, ServeGlobal: ServeStateType }>();
        const data: any = ref(null);
        const expandIt = (tree: any, options: any, e: any) => {
            const {parent, ancestor, isRoot} = options;
            // 异步获取组件详情信息
            if (tree.ref) {
                // 如果没有引用组件内容，需要获取组件详情
                if (!tree.content) {
                    const result: any = (props.components || []).find((item: any) => item.ref === tree.ref);
                    tree.content = JSON.parse(result.content || '{}');
                    data.value = addExtraViewInfo(data.value);
                } else {
                    delete tree.content;
                }
            } else {
                tree.extraViewInfo.isExpand = !tree.extraViewInfo.isExpand;
            }
        }

        watch(() => {
            return props.value
        }, (newVal) => {
            const val = cloneDeep(newVal);
            data.value = addExtraViewInfo(val);
        }, {
            immediate: true,
            deep: true
        });
        watch(() => {
            return data.value
        }, (newVal) => {
            const newObj = removeExtraViewInfo(cloneDeep(newVal));
            emit('change', newObj);
        }, {
            immediate: true,
            deep: true
        });
        const renderDivider = (options: any) => {
            const {isRoot, isFirst, isLast, keyIndex, parent, ancestor, isRefChildNode, keyName} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            const required = (items?.required || []).includes(keyName);
            if (!required) {
                return null;
            }
            return <div class={'action'}>
                <SplitDivider/>
            </div>
        }
        const renderExtraAction = (options: any) => {
            const {isRoot, keyIndex, keyName, parent, tree, ancestor, isRefChildNode} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            const required = (items?.required || []).includes(keyName);
            return <div class={'extraAction'}>
                <ExtraInfo
                    isRoot={isRoot}
                    required={required}
                    deprecated={tree?.deprecated}
                    isRefChildNode={isRefChildNode || false}
                    value={tree}/>
            </div>
        }
        const renderDataTypeSetting = (options: any) => {
            const {tree, isRefChildNode} = options;
            const propsLen = Object.keys(tree?.properties || {}).length;
            return <>
                <DataTypeSetting refsOptions={props.refsOptions}
                                 value={tree}
                                 isRefChildNode={isRefChildNode || false}/>
                {isObject(tree?.type) && !isRef(tree) ? <span
                    class={'baseInfoSpace'}>{tree.types?.length > 0 ? `[${propsLen}]` : `{${propsLen}}`}</span> : null}
            </>
        }

        const renderProperties = (options: any) => {
            const {keyName, parent, depth, isRoot} = options;
            if (isRoot) {
                return null
            }
            if (!parent) {
                return null;
            }

            const properties = parent?.properties?.[keyName] || {};
            console.log('832 properties', properties)
            const list: any = [];
            Object.entries(properties).forEach(([k, v]) => {
                if (typeof v !== 'boolean' && !['type', 'properties', 'extraViewInfo','ref','content','name'].includes(k)) {
                    if (!!v || v === 0) {
                        list.push({
                            label: k,
                            value: v
                        })
                    }
                }
            })
            if (list.length === 0) {
                return null
            }
            return <div>
                {
                    list.map((item) => {
                        const {label, value} = item;
                        return <div class={['directoryText', 'properties-info']}
                                    style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                            {label !== 'description' ?
                                <a-typography-text type="secondary">{label}：</a-typography-text> : null}
                            {label === 'description' ? <a-typography-text>{value}</a-typography-text> : null}
                            {label !== 'description' ?
                                <a-typography-text type="secondary">{value}</a-typography-text> : null}
                        </div>
                    })
                }
            </div>

        }
        const renderKeyName = (options: any) => {
            const {keyName, keyIndex, parent, isRoot, isRefChildNode, isRef, isRefRootNode} = options;
            if (isRoot) return null;
            if (!keyName) return null;
            if (isRefRootNode) return null;
            return <>
                <span class={'baseInfoKey'}>
                    {keyName}
                </span>
                <span class={'baseInfoSpace'}>:</span>
            </>
        }
        const renderExpandIcon = (options: any) => {
            const {isExpand, tree, isRoot, isRef} = options;
            // if (isRoot) return null;
            if (isExpand) {
                return <DownOutlined onClick={expandIt.bind(this, tree, options)} class={'expandIcon'}/>
            } else {
                return <RightOutlined onClick={expandIt.bind(this, tree, options)} class={'expandIcon'}/>
            }
        }
        const renderHorizontalLine = (depth: number) => {
            if (depth === 1) return null;
            return <div class={'horizontalLine'}
                        style={{left: `${(depth - 1) * treeLevelWidth + 8}px`}}/>
        }
        const renderNormalType = (options: any) => {
            const {isExpand} = options;
            return (<div>
                <div key={options.index}
                     class={'leafNode'} style={{'paddingLeft': `${options.depth * treeLevelWidth}px`}}>
                    {!options.isRoot ? <div class={'leafNodeHorizontalLine'}
                                            style={{left: `${(options.depth - 1) * treeLevelWidth + 8}px`}}/> : null}
                    <div class={'baseInfo'}>
                        {renderKeyName(options)}
                        {renderDataTypeSetting(options)}
                    </div>
                    {renderDivider(options)}
                    {renderExtraAction(options)}
                </div>
                {isExpand ? renderProperties(options) : null}
            </div>)
        }
        const renderDirectoryText = (options: any) => {
            const {depth, tree, isRefChildNode, isRoot, isExpand} = options;
            return <div class={'helo'}>
                <div class={'directoryText'}
                     style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                    {renderHorizontalLine(depth)}
                    <div class={'baseInfo'}>
                        {renderExpandIcon(options)}
                        {renderKeyName(options)}
                        {renderDataTypeSetting(options)}
                    </div>
                    {renderDivider(options)}
                    {renderExtraAction(options)}
                </div>
                {isExpand ? renderProperties(options) : null}
            </div>
        }

        const renderVerticalLine = (options: any) => {
            return <div class={'verticalLine'} style={{left: `${options.depth * treeLevelWidth + 8}px`}}></div>
        }
        const renderTree = (tree: any) => {
            if (!tree) return null;
            const isRoot = tree?.extraViewInfo?.depth === 1;
            const options = {...tree?.extraViewInfo, isRoot, tree: tree};
            if (isNormalType(tree.type) && !isRef(tree)) {
                return renderNormalType(options)
            }
            // 渲染对象类型节点
            if (isObject(tree.type) && !isRef(tree)) {
                const isRoot = tree?.extraViewInfo?.depth === 1;
                const isExpand = tree?.extraViewInfo?.isExpand;
                const options = {...tree?.extraViewInfo, isRoot, tree}
                return <div key={tree.type} class={{'directoryNode': true, "rootNode": isRoot}}>
                    {renderDirectoryText(options)}
                    {
                        isExpand && Object.entries(tree?.properties || {}).map(([key, value]: any) => {
                            return renderTree(value)
                        })
                    }
                    {isExpand && Object.keys(tree?.properties || {}).length > 0 && renderVerticalLine(options)}
                </div>
            }
            // 渲染数组类型节点
            if (isArray(tree.type) && !isRef(tree)) {
                // 找到最后一个非数组类型的节点
                const {node} = findLastNotArrayNode(tree);
                const isRoot = tree?.extraViewInfo?.depth === 1;
                return <div class={{'directoryNode': true, "rootNode": isRoot}}>
                    {
                        renderTree(node)
                    }
                </div>
            }
            // 如果是引用类型
            if (isRef(tree)) {
                const isRoot = tree?.extraViewInfo?.depth === 1;
                const isExpand = tree?.extraViewInfo?.isExpand;
                const isRef = tree?.extraViewInfo?.isRef;
                const options = {...tree?.extraViewInfo, isRoot, tree}
                return <div key={tree.type}
                            class={{'directoryNode': true, "rootNode": isRoot, 'refNode': isRef || !!tree?.ref}}>
                    {renderDirectoryText(options)}
                    {
                        isExpand && tree?.content && renderTree(tree.content)
                    }
                    {isExpand && renderVerticalLine(options)}
                </div>
            }
        }
        return () => (
            <div class={'schemaViewer-content'} style={props.contentStyle}>
                {renderTree(data.value)}
            </div>
        )
    }
})
