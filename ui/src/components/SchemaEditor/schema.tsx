import {defineComponent, ref, watch,} from 'vue';
import './schema.less';
import {DownOutlined, PlusOutlined, RightOutlined,} from '@ant-design/icons-vue';
import Actions from "./Actions.vue";
import ExtraActions from "./ExtraActions.vue";
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
        refsOptions: Array
    },
    emits: ['change'],
    setup(props, {emit}) {
        const store = useStore<{ Endpoint, ServeGlobal: ServeStateType }>();
        const data: any = ref(null);
        const expandIt = async (tree: any, options: any, e: any) => {
            const {parent, ancestor, isRoot} = options;
            // 异步获取组件详情信息
            if (tree.ref) {
                // 如果没有引用组件内容，需要获取组件详情
                if (!tree.content) {
                    const result = await store.dispatch('Endpoint/getRefDetail', {
                        ref: tree.ref,
                        serveId: props.serveId
                    })
                    tree.content = JSON.parse(result.content || '{}');
                    data.value = addExtraViewInfo(data.value);
                } else {
                    delete tree.content;
                }
            } else {
                tree.extraViewInfo.isExpand = !tree.extraViewInfo.isExpand;
            }
        }
        const addProps = (tree: any, e: any) => {
            tree.properties = tree.properties || {};
            const keys = Object.keys(tree.properties);
            keys.push(`name${keys.length + 1}`);
            const newVal: any = {type: 'string'};
            const newObj: any = {};
            keys.forEach((item) => {
                if (tree.properties[item]) {
                    newObj[item] = tree.properties[item];
                } else {
                    newObj[item] = newVal;
                }
            })
            tree.properties = {...newObj};
            data.value = addExtraViewInfo(data.value);
        }
        const pasteKeyName = (e) => {
            // 阻止默认的粘贴事件
            e.preventDefault();
            // 从剪贴板中获取纯文本
            const text = (e.originalEvent || e).clipboardData.getData('text/plain');
            // 插入纯文本
            document.execCommand("insertHTML", false, text);
        }
        const updateKeyName = (oldKey: any, keyIndex: any, parent: any, event: any) => {
            const newKey = event.target.innerText;
            const keys = Object.keys(parent.properties);
            // 判断组件名称必须是英文，复合 URL 规则
            // 匹配由 数字、26个英文字母、下划线、 - 组成的字符串
            // const reg = /^[a-zA-Z$][\w\W]*$/;
            // if(!reg.test(newKey)){
            //     message.warning(`属性须以字母开头`);
            // }
            // 新旧 key 相等
            if (oldKey === newKey) {
                return;
                //  已经存在了 key
            } else if (keys.includes(newKey)) {
                message.warning(`已存在名为${newKey}的属性`);
                keys[keyIndex] = oldKey;
            } else if (newKey) {
                keys[keyIndex] = newKey;
            } else if (!newKey) {
                message.warning(`属性名不能为空`);
                keys[keyIndex] = oldKey;
            }
            const newObj: any = {};
            keys.forEach((item) => {
                newObj[item] = parent.properties[item] || parent.properties[oldKey];
            })
            parent.properties = {...newObj};
            data.value = addExtraViewInfo(data.value);
        }
        const dataTypeChange = (options?: any, newProps?: any) => {
            const {parent, keyName, depth, ancestor} = options;
            const firstType = newProps?.[0]?.type;
            // 如果是根节点
            if (depth === 1) {
                if (isArray(firstType)) {
                    data.value = generateSchemaByArray(newProps);
                } else {
                    data.value = {
                        ...newProps[0]
                    }
                }
                // 非根节点
            } else {
                //  数组类型, 且个数大于等于 1 ，需要生成新的schema
                if (newProps?.length >= 1 && isArray(firstType)) {
                    const items = parent?.type === 'array' ? ancestor : parent;
                    if (items?.properties?.[keyName]) {
                        items.properties[keyName] = generateSchemaByArray(newProps);
                        console.log('123', items.properties[keyName]);
                    }
                    // 非数组类型
                } else if (newProps?.length === 1 && !isArray(firstType)) {
                    const items = parent?.type === 'array' ? ancestor : parent;
                    if (items?.properties?.[keyName]) {
                        items.properties[keyName] = Object.assign(items.properties[keyName], {...newProps[0]})
                    }
                }
            }
            data.value = addExtraViewInfo(data.value);
            console.log('change datatype  data.value', data.value);
        }
        const moveUp = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            // 互换两个元素的位置
            [keys[keyIndex - 1], keys[keyIndex]] = [keys[keyIndex], keys[keyIndex - 1]];
            const newObj: any = {};
            keys.forEach((item) => {
                newObj[item] = parent.properties[item];
            })
            parent.properties = {...newObj};
            data.value = addExtraViewInfo(data.value);
        };
        const moveDown = (keyIndex: number, parent: any) => {
            const keys = Object.keys(parent.properties);
            // 互换两个元素的位置
            [keys[keyIndex + 1], keys[keyIndex]] = [keys[keyIndex], keys[keyIndex + 1]];
            const newObj: any = {};
            keys.forEach((item) => {
                newObj[item] = parent.properties[item];
            })
            parent.properties = {...newObj};
            data.value = addExtraViewInfo(data.value);
        };
        const copy = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            const key = keys[keyIndex];
            const copyObj = cloneDeep(parent.properties[key]);
            let keyCopyName = `${key}-copy`;
            if (keys.includes(keyCopyName)) {
                keyCopyName = `${keyCopyName}-copy`;
            }
            keys.splice(keyIndex + 1, 0, `${keyCopyName}`);
            console.log('keys', keys);
            const newObj: any = {};
            keys.forEach((item, index: number) => {
                if (parent.properties[item]) {
                    newObj[item] = parent.properties[item];
                } else {
                    newObj[item] = copyObj;
                }
            })
            parent.properties = {...newObj};
            data.value = addExtraViewInfo(data.value);
        }
        const setRequire = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            const key = keys[keyIndex];
            // ::::todo 有问题,不能选择
            if (!parent?.required?.includes(key)) {
                if (!parent.required) {
                    parent.required = [key];
                } else {
                    parent.required.push(key);
                }
            }
        };
        const addDesc = (tree: any, desc: string) => {
            tree.description = desc;
            data.value = addExtraViewInfo(data.value);
        };
        const del = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            keys.splice(keyIndex, 1);
            const newObj: any = {};
            keys.forEach((item) => {
                newObj[item] = parent.properties[item];
            })
            parent.properties = {...newObj};
            data.value = addExtraViewInfo(data.value);
        };
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
        const renderAction = (options: any) => {
            const {isRoot, isFirst, isLast, keyIndex, parent, ancestor, isRefChildNode} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            return <div class={'action'}>
                <Actions
                    isRoot={isRoot}
                    isFirst={isFirst || false}
                    isLast={isLast || false}
                    isRefChildNode={isRefChildNode || false}
                    onMoveDown={moveDown.bind(this, keyIndex, items)}
                    onMoveUp={moveUp.bind(this, keyIndex, items)}
                    onCopy={copy.bind(this, keyIndex, items)}/>
            </div>
        }
        const renderExtraAction = (options: any) => {
            const {isRoot, keyIndex, parent, tree, ancestor, isRefChildNode} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            return <div class={'extraAction'}>
                <ExtraActions
                    isRoot={isRoot}
                    isRefChildNode={isRefChildNode || false}
                    value={tree}
                    onAddDesc={addDesc.bind(this, tree)}
                    onDel={del.bind(this, keyIndex, items)}
                    onSetRequire={setRequire.bind(this, keyIndex, items)}/>
            </div>
        }
        const renderDataTypeSetting = (options: any) => {
            const {tree, isRefChildNode} = options;
            const propsLen = Object.keys(tree?.properties || {}).length;
            return <>
                <DataTypeSetting refsOptions={props.refsOptions}
                                 value={tree}
                                 isRefChildNode={isRefChildNode || false}
                                 onChange={dataTypeChange.bind(this, options)}/>
                {isObject(tree?.type) && !isRef(tree) ? <span
                    class={'baseInfoSpace'}>{tree.types?.length > 0 ? `[${propsLen}]` : `{${propsLen}}`}</span> : null}
            </>
        }
        const renderKeyName = (options: any) => {
            const {keyName, keyIndex, parent, isRoot, isRefChildNode, isRef, isRefRootNode} = options;
            if (isRoot) return null;
            if (!keyName) return null;
            if (isRefRootNode) return null;
            return <>
                <span class={'baseInfoKey'}
                      contenteditable={!isRefChildNode}
                      onPaste={pasteKeyName}
                      onBlur={updateKeyName.bind(this, keyName, keyIndex, parent)}>
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
            return (<div key={options.index}
                         class={'leafNode'} style={{'paddingLeft': `${options.depth * treeLevelWidth}px`}}>
                {!options.isRoot ? <div class={'leafNodeHorizontalLine'}
                                        style={{left: `${(options.depth - 1) * treeLevelWidth + 8}px`}}/> : null}
                <div class={'baseInfo'}>
                    {renderKeyName(options)}
                    {renderDataTypeSetting(options)}
                </div>
                {renderAction(options)}
                {renderExtraAction(options)}
            </div>)
        }
        const renderDirectoryText = (options: any) => {
            const {depth, tree, isRefChildNode} = options;
            return <div class={'directoryText'}
                        style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                {renderHorizontalLine(depth)}
                <div class={'baseInfo'}>
                    {renderExpandIcon(options)}
                    {renderKeyName(options)}
                    {renderDataTypeSetting(options)}
                    {isObject(tree.type) && !isRef(tree) && !isRefChildNode ?
                        <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/> : null}
                </div>
                {renderAction(options)}
                {renderExtraAction(options)}
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
            <div class={'schemaEditor-content'} style={props.contentStyle}>
                {renderTree(data.value)}
            </div>
        )
    }
})
