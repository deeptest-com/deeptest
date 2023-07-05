import {defineComponent, ref, watch,nextTick} from 'vue';
import './schema.less';
import {DownOutlined, PlusOutlined, RightOutlined,} from '@ant-design/icons-vue';
import SplitDivider from "./SplitDivider.vue";
import ExtraInfo from "./ExtraInfo.vue";
import DataTypeSetting from './DataTypeSetting.vue';
import cloneDeep from "lodash/cloneDeep";
import {
    addExtraViewInfo,
    findLastNotArrayNode, handleRefInfo,
    isArray,
    isNormalType,
    isObject,
    isRef,
} from '@/components/SchemaEditor/utils';
import {
    treeLevelWidth
} from '@/components/SchemaEditor/config';

export default defineComponent({
    name: 'SchemeEditor',
    props: {
        value: Object,
        contentStyle: Object,
        serveId: Number,
        refsOptions: Array,
        components: Array,
    },
    emits: [],
    setup(props, {emit}) {
        const data: any = ref(null);
        const expandIt = (tree: any, options: any, e: any) => {
            const {parent, ancestor, isRoot} = options;
            // 异步获取组件详情信息
            if (tree.ref) {
                // 如果没有引用组件内容，需要获取组件详情
                if (!tree.content) {
                    const result: any = (props.components || []).find((item: any) => item.ref === tree.ref);
                    // 处理引用组件的信息
                    handleRefInfo(tree, result);
                    data.value = addExtraViewInfo(data.value);
                    tree.extraViewInfo.isExpand = true;


                } else {
                    delete tree.content;
                    tree.extraViewInfo.isExpand = false;
                }
            } else {
                tree.extraViewInfo.isExpand = !tree.extraViewInfo.isExpand;
            }
        }

        watch(() => {return props.value}, (newVal) => {
            data.value = addExtraViewInfo(newVal);
        }, {immediate: true, deep: false});

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
            const list: any = [];
            Object.entries(properties).forEach(([k, v]) => {
                if (typeof v !== 'boolean' && !['type', 'properties', 'extraViewInfo','ref','content','name','required'].includes(k)) {
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
            return <div>
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
                return <div class={{'directoryNode': true, "rootNode": isRoot,'rootNode-array':isRoot}}>
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
