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
    isObject, moveCursorToEnd,
    removeExtraViewInfo,
    treeLevelWidth
} from './utils';
import {message} from "ant-design-vue";

export default defineComponent({
    name: 'SchemeEditor',
    props: {
        value: Object,
        contentStyle: Object,
        refsOptions: Array
    },
    emits: ['change'],
    setup(props, {emit}) {
        const data: any = ref(null);
        const expandIt = (tree: any, e: any) => {
            if (tree?.extraViewInfo) {
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
            }else if(newKey){
                keys[keyIndex] = newKey;
            }else if(!newKey){
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
            if (parent === null && depth === 1) {
                if (isArray(firstType)) {
                    data.value = generateSchemaByArray(newProps);
                } else {
                    data.value = {
                        ...newProps[0]
                    }
                }
                // 非根节点
            } else {
                //  数组类型, 且个数大于 1 ，需要生成新的schema
                if (newProps?.length > 1 && isArray(firstType)) {
                    const items = parent?.type === 'array' ? ancestor : parent;
                    if (items?.properties?.[keyName]) {
                        items.properties[keyName] = generateSchemaByArray(newProps);
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
            if (!parent.required.includes(key)) {
                parent.required.push(key);
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
            const {isRoot, isFirst, isLast, keyIndex, parent, ancestor} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            return <div class={'action'}>
                <Actions
                    isRoot={isRoot}
                    isFirst={isFirst || false}
                    isLast={isLast || false}
                    onMoveDown={moveDown.bind(this, keyIndex, items)}
                    onMoveUp={moveUp.bind(this, keyIndex, items)}
                    onCopy={copy.bind(this, keyIndex, items)}/>
            </div>
        }
        const renderExtraAction = (options: any) => {
            const {isRoot, keyIndex, parent, tree, ancestor} = options;
            const items = parent?.type === 'array' ? ancestor : parent;
            return <div class={'extraAction'}>
                <ExtraActions
                    isRoot={isRoot}
                    value={tree}
                    onAddDesc={addDesc.bind(this, tree)}
                    onDel={del.bind(this, keyIndex, items)}
                    onSetRequire={setRequire.bind(this, keyIndex, items)}/>
            </div>
        }
        const renderDataTypeSetting = (options: any) => {
            const {tree} = options;
            const propsLen = Object.keys(tree?.properties || {}).length;
            return <>
                <DataTypeSetting refsOptions={props.refsOptions}
                                 value={tree}
                                 onChange={dataTypeChange.bind(this, options)}/>
                {isObject(tree?.type) ? <span
                    class={'baseInfoSpace'}>{tree.types?.length > 0 ? `[${propsLen}]` : `{${propsLen}}`}</span> : null}
            </>
        }
        const renderKeyName = (options: any) => {
            const {keyName, keyIndex, parent, isRoot} = options;
            if (isRoot) return null;
            return <>
                <span class={'baseInfoKey'}
                      contenteditable={true}
                      onBlur={updateKeyName.bind(this, keyName, keyIndex, parent)}>
                    {keyName}
                </span>
                <span class={'baseInfoSpace'}>:</span>
            </>
        }
        const renderExpandIcon = (options: any) => {
            const {isExpand, tree, isRoot} = options;
            if (isRoot) return null;
            if (isExpand) {
                return <DownOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/>
            } else {
                return <RightOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/>
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
            const {depth, tree} = options;
            return <div class={'directoryText'}
                        style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                {renderHorizontalLine(depth)}
                <div class={'baseInfo'}>
                    {renderExpandIcon(options)}
                    {renderKeyName(options)}
                    {renderDataTypeSetting(options)}
                    {isObject(tree.type) ? <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/> : null}
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
            if (isNormalType(tree.type)) {
                return renderNormalType(options)
            }
            // 渲染对象类型节点
            if (isObject(tree.type)) {
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
            if (isArray(tree.type)) {
                // 找到最后一个非数组类型的节点
                const {node} = findLastNotArrayNode(tree);
                const isRoot = tree?.extraViewInfo?.depth === 1;
                return <div class={{'directoryNode': true, "rootNode": isRoot}}>
                    {
                        renderTree(node)
                    }
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
