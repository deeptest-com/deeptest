import {defineComponent, ref, watch, nextTick} from 'vue';
import './schema.less';

import {DownOutlined, PlusOutlined, RightOutlined,} from '@ant-design/icons-vue';
import Actions from "./Actions.vue";
import ExtraActions from "./ExtraActions.vue";
import DataTypeSetting from './DataTypeSetting.vue';
import MockRule from "./MockRule.vue";
import cloneDeep from "lodash/cloneDeep";
import {
    addExtraViewInfo,
    findLastNotArrayNode,
    generateSchemaByArray,
    handleRefInfo,
    isArray,
    isCompositeType,
    isNormalType,
    isObject,
    isRef,
    removeExtraViewInfo,
} from './utils';
import {
    treeLevelWidth
} from './config';
import {message, notification} from "ant-design-vue";
import {useStore} from "vuex";
import {StateType as ServeStateType} from "@/store/serve";
import {notifyWarn} from "@/utils/notify";



export default defineComponent({
    name: 'SchemeEditor',
    props: ['value', 'contentStyle', 'serveId'],
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
                    // 处理引用组件的信息
                    handleRefInfo(tree, result);
                    data.value = addExtraViewInfo(data.value);
                    tree.extraViewInfo.isExpand = true;
                } else {
                    tree.extraViewInfo.isExpand = false;
                    delete tree.content;
                }
            } else {
                tree.extraViewInfo.isExpand = !tree.extraViewInfo.isExpand;
            }
        }
        const addProps = (tree: any, e: any) => {
            // 如果是对象类型，添加属性
            if (isObject(tree.type)) {
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
            }
            if (isCompositeType(tree.type)) {
                // 如果复合类型是数组，添加元素
                if (Array.isArray(tree[tree.type])) {
                    tree[tree.type].push({type: 'string'});
                    //  否则添加属性
                } else {
                    tree[tree.type] = [{type: 'string'}]
                }
            }
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

        const updateMockType = (tree: any, mockType: string) => {
            tree.mockType = mockType;
            data.value = addExtraViewInfo(data.value);
        };

        const keyNameKeyDown = (oldKey: any, keyIndex: any, parent: any, event: any) => {
            // 获取用户输入的字符
            const char = event.key;
            console.log('832char', char);
            // 允许输入字母、数字、下划线、短横线和  删除、上下左右箭头键
            if (/^[\w-]$/.test(char) || ['Delete', 'Backspace', 'ArrowLeft', 'ArrowRight', 'ArrowUp', 'ArrowDown'].includes(char)) {
                console.log('合法字符');
            } else {
                event.preventDefault();
            }
        }
        const updateKeyName = (oldKey: any, keyIndex: any, parent: any, event: any) => {
            const newKey = event.target.innerText;
            const keys = Object.keys(parent.properties);
            // const reg = /^\w+$/;
            // if(!reg.test(newKey)){
            //     notifyWarn(`属性名非法，请重新输入`);
            //     event.target.innerText = oldKey;
            //     event.preventDefault();
            //     return;
            // }
            // 新旧 key 相等
            if (oldKey === newKey) {
                return;
                //  已经存在了 key
            } else if (keys.includes(newKey)) {
                notifyWarn(`已存在名为${newKey}的属性`);
                keys[keyIndex] = oldKey;
            } else if (newKey) {
                keys[keyIndex] = newKey;
            } else if (!newKey) {
                notifyWarn(`属性名不能为空`);
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
            if (!firstType) {
                return;
            }
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
                    // 如果是复合类型，需要特殊处理
                    if (isCompositeType(ancestor?.type)) {
                        if (ancestor?.[ancestor.type]?.[keyName]) {
                            ancestor[ancestor.type][keyName] = generateSchemaByArray(newProps);
                        }
                    } else {
                        const items = parent?.type === 'array' ? ancestor : parent;
                        if (items?.properties?.[keyName]) {
                            items.properties[keyName] = generateSchemaByArray(newProps);
                        }
                    }
                    // 非数组类型
                } else if (newProps?.length >= 1 && !isArray(firstType)) {
                    // 如果是复合类型，需要特殊处理
                    if (isCompositeType(ancestor?.type)) {
                        if (ancestor?.[ancestor.type]?.[keyName]) {
                            ancestor[ancestor.type][keyName] = {...newProps[0]};
                        }
                    } else {
                        const items = parent?.type === 'array' ? ancestor : parent;
                        if (items?.properties?.[keyName]) {
                            // 既然更换类型，之前的属性就不需要了
                            items.properties[keyName] = {...newProps[0]};
                        }
                    }
                // 其他场景
                }  else {
                    notifyWarn(`未知异常，请重试`);
                }
            }
            data.value = addExtraViewInfo(data.value);
            console.log('change datatype  data.value', data.value);
        }
        const moveUp = (keyIndex: any, parent: any) => {
            if (isCompositeType(parent.type)) {
                const combines = {
                    allOf: parent?.allOf || [],
                    oneOf: parent?.oneOf || [],
                    anyOf: parent?.anyOf || [],
                }
                const items = combines[parent.type];
                // 互换两个元素的位置
                [items[keyIndex - 1], items[keyIndex]] = [items[keyIndex], items[keyIndex - 1]];
                parent[parent.type] = [...items];
                data.value = addExtraViewInfo(data.value);
            } else {
                const keys = Object.keys(parent.properties);
                // 互换两个元素的位置
                [keys[keyIndex - 1], keys[keyIndex]] = [keys[keyIndex], keys[keyIndex - 1]];
                const newObj: any = {};
                keys.forEach((item) => {
                    newObj[item] = parent.properties[item];
                })
                parent.properties = {...newObj};
                data.value = addExtraViewInfo(data.value);
            }
        };
        const moveDown = (keyIndex: number, parent: any) => {
            if (isCompositeType(parent.type)) {
                const combines: any = {
                    allOf: parent?.allOf || [],
                    oneOf: parent?.oneOf || [],
                    anyOf: parent?.anyOf || [],
                }
                const items = combines[parent.type];
                // 互换两个元素的位置
                [items[keyIndex + 1], items[keyIndex]] = [items[keyIndex], items[keyIndex + 1]];
                parent[parent.type] = [...items];
                data.value = addExtraViewInfo(data.value);

            } else {
                const keys = Object.keys(parent.properties);
                // 互换两个元素的位置
                [keys[keyIndex + 1], keys[keyIndex]] = [keys[keyIndex], keys[keyIndex + 1]];
                const newObj: any = {};
                keys.forEach((item) => {
                    newObj[item] = parent.properties[item];
                })
                parent.properties = {...newObj};
                data.value = addExtraViewInfo(data.value);
            }
        };
        const copy = (keyIndex: any, parent: any) => {
            if (isCompositeType(parent.type)) {
                const combines = {
                    allOf: parent?.allOf || [],
                    oneOf: parent?.oneOf || [],
                    anyOf: parent?.anyOf || [],
                }
                const copyObj = cloneDeep(combines[parent.type][keyIndex]);
                combines[parent.type].splice(keyIndex + 1, 0, copyObj);
                parent[parent.type] = combines[parent.type];
                data.value = addExtraViewInfo(data.value);
            } else {
                const keys = Object.keys(parent.properties);
                const key = keys[keyIndex];
                const copyObj = cloneDeep(parent.properties[key]);
                let keyCopyName = `${key}-copy`;
                if (keys.includes(keyCopyName)) {
                    keyCopyName = `${keyCopyName}-copy`;
                }
                keys.splice(keyIndex + 1, 0, `${keyCopyName}`);
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
        }
        const setRequire = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            const key = keys[keyIndex];
            parent.required = Array.isArray(parent?.required) ? parent?.required : [];
            if (!parent.required.includes(key)) {
                parent.required.push(key);
            } else {
                const index = parent.required.indexOf(key);
                parent.required.splice(index, 1);
            }
        };

        const addDesc = (tree: any, desc: string) => {
            tree.description = desc;
            data.value = addExtraViewInfo(data.value);
        };

        const del = (keyIndex: any, parent: any) => {
            if (isCompositeType(parent.type)) {
                const combines = {
                    allOf: parent?.allOf || [],
                    oneOf: parent?.oneOf || [],
                    anyOf: parent?.anyOf || [],
                }
                combines[parent.type].splice(keyIndex, 1);
                parent[parent.type] = combines[parent.type];
                data.value = addExtraViewInfo(data.value);
            } else {
                const keys = Object.keys(parent.properties);
                keys.splice(keyIndex, 1);
                const newObj: any = {};
                keys.forEach((item) => {
                    newObj[item] = parent.properties[item];
                })
                parent.properties = {...newObj};
                data.value = addExtraViewInfo(data.value);
            }
        };

        // 监听value变化，更新data，他是一个字符串
        watch(() => {
            return props.value
        }, (newVal: any) => {
            try {
                nextTick(() => {
                    let obj = JSON.parse(newVal);
                    obj = obj ? obj : {type: 'object'};
                    data.value = addExtraViewInfo(obj);
                })
            } catch (e) {
                console.log('watch', e);
            }
        }, {immediate: true});

        watch(() => {
            return data.value
        }, (newVal) => {
            const newObj = removeExtraViewInfo(cloneDeep(newVal), true);
            emit('change', newObj);
        }, {
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
            const {isRoot, keyIndex, parent, tree, ancestor, isRefChildNode, isCompositeChildNode} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            return <div class={'extraAction'}>
                <ExtraActions
                    isRoot={isRoot}
                    isRefChildNode={isRefChildNode || false}
                    isCompositeChildNode={isCompositeChildNode || false}
                    value={tree}
                    onAddDesc={addDesc.bind(this, tree)}
                    onDel={del.bind(this, keyIndex, items)}
                    onSetRequire={setRequire.bind(this, keyIndex, items)}/>
            </div>
        }
        const renderDataTypeSetting = (options: any) => {
            const {tree, isRefChildNode, isRoot} = options;
            const propsLen = Object.keys(tree?.properties || {}).length;
            const combines = {
                allOf: tree?.allOf || [],
                oneOf: tree?.oneOf || [],
                anyOf: tree?.anyOf || [],
            }
            return <>
                <DataTypeSetting
                    value={tree}
                    serveId={props.serveId}
                    isRefChildNode={isRefChildNode || false}
                    isRoot={isRoot || false}
                    onChange={dataTypeChange.bind(this, options)}/>
                {isObject(tree?.type) && !isRef(tree) ? <span
                    class={'baseInfoSpace'}>{tree.types?.length > 0 ? `[${propsLen}]` : `{${propsLen}}`}</span> : null}
                {/* 复合类型 */}
                {isCompositeType(tree?.type) ?
                    <span class={'baseInfoSpace'}>{`{${combines[tree.type].length}}`}</span> : null}
            </>
        }
        const renderKeyName = (options: any) => {
            const {
                keyName,
                keyIndex,
                parent,
                isRoot,
                isRefChildNode,
                isRef,
                isRefRootNode,
                ancestor,
                isCompositeChildNode
            } = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            if (isRoot) return null;
            if (!keyName) return null;
            if (isRefRootNode || isCompositeChildNode) return null;
            return <>
                <span class={'baseInfoKey'}
                      contenteditable={!isRefChildNode}
                      onPaste={pasteKeyName}
                      onKeydown={keyNameKeyDown.bind(this, keyName, keyIndex, items)}
                      onBlur={updateKeyName.bind(this, keyName, keyIndex, items)}>
                    {keyName}
                </span>
                <span class={'baseInfoSpace'}>:</span>
            </>
        }
        const renderExpandIcon = (options: any) => {
            const {isExpand, tree, isRoot, isRef} = options;
            if (isExpand) {
                return <DownOutlined onClick={expandIt.bind(this, tree, options)} class={'expandIcon'}/>
            } else {
                return <RightOutlined onClick={expandIt.bind(this, tree, options)} class={'expandIcon'}/>
            }
        }

        const renderMockRule = (options: any) => {
            const {tree, isRefChildNode, isRoot} = options;
            return <MockRule  onUpdate={updateMockType.bind(this,tree)}/>
        }

        const renderHorizontalLine = (depth: number) => {
            if (depth === 1) return null;
            return <div class={'horizontalLine'}
                        style={{left: `${(depth - 1) * treeLevelWidth + 8}px`}}/>
        }
        const renderNormalType = (options: any) => {
            const { tree} = options;
            return (<div key={options.index}
                         class={'leafNode'} style={{'paddingLeft': `${options.depth * treeLevelWidth}px`}}>
                {!options.isRoot ? <div class={'leafNodeHorizontalLine'}
                                        style={{left: `${(options.depth - 1) * treeLevelWidth + 8}px`}}/> : null}
                <div class={'baseInfo'}>
                    {renderKeyName(options)}
                    {renderDataTypeSetting(options)}
                    {!isRef(tree) ? renderMockRule(options) : null}
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
                    {
                        (isObject(tree.type) || isCompositeType(tree.type)) && !isRef(tree) && !isRefChildNode ?
                            <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/> : null
                    }
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

            // 普通类型
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
            // 渲染复合类型节点
            if (isCompositeType(tree.type) && !isRef(tree)) {
                const isRoot = tree?.extraViewInfo?.depth === 1;
                const isExpand = tree?.extraViewInfo?.isExpand;
                const options = {...tree?.extraViewInfo, isRoot, tree}

                const combines = {
                    allOf: tree?.allOf || [],
                    oneOf: tree?.oneOf || [],
                    anyOf: tree?.anyOf || [],
                }
                return <div key={tree.type} class={{'directoryNode': true, "rootNode": isRoot}}>
                    {renderDirectoryText(options)}
                    {
                        isExpand && combines[tree.type].map((value: any) => {
                            return renderTree(value)
                        })
                    }
                    {isExpand && combines[tree.type].length > 0 && renderVerticalLine(options)}
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

        return () => {
            if (!data.value) return null;
            return (
                <div class={'schemaEditor-content'} style={props.contentStyle}>
                    {renderTree(data.value)}
                </div>
            )
        }
    }
})
