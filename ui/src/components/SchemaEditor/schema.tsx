import {
    defineComponent,
    ref,
    watch,
} from 'vue';
import './schema.less';
import {
    DownOutlined,
    RightOutlined,
    PlusOutlined,
} from '@ant-design/icons-vue';
import Actions from "./Actions.vue";
import ExtraActions from "./ExtraActions.vue";
import DataTypeSetting from './DataTypeSetting.vue';
import cloneDeep from "lodash/cloneDeep";
import {removeExtraInfo, addExtraInfo, isArray, isObject,treeLevelWidth} from './utils';

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
            data.value = addExtraInfo(data.value);
        }
        const updateKeyName = (oldKey, keyIndex, parent, event) => {
            const newKey = event.target.innerText;
            const keys = Object.keys(parent.properties);
            keys[keyIndex] = newKey;
            const newObj: any = {};
            keys.forEach((item) => {
                if (item === newKey) {
                    newObj[item] = parent.properties[oldKey];
                } else {
                    newObj[item] = parent.properties[item];
                }
            })
            parent.properties = {...newObj};
        }
        const dataTypeChange = (tree: any, newVal) => {
            const [key, value]: any = Object.entries(newVal)[0];
            if (key) {
                tree.type = key;
                tree = Object.assign(tree, value);
            }
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
        };
        const moveDown = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            // 互换两个元素的位置
            [keys[keyIndex + 1], keys[keyIndex]] = [keys[keyIndex], keys[keyIndex + 1]];
            const newObj: any = {};
            keys.forEach((item) => {
                newObj[item] = parent.properties[item];
            })
            parent.properties = {...newObj};
        };
        const copy = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            const key = keys[keyIndex];
            const copyObj = JSON.parse(JSON.stringify(parent.properties[key]));
            keys.splice(keyIndex + 1, 0, `${key}-copy`);
            const newObj: any = {};
            keys.forEach((item) => {
                if (parent.properties[item]) {
                    newObj[item] = parent.properties[item];
                } else {
                    newObj[item] = copyObj;
                }
            })
            parent.properties = {...newObj};
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
        };
        const del = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            keys.splice(keyIndex, 1);
            const newObj: any = {};
            keys.forEach((item) => {
                newObj[item] = parent.properties[item];
            })
            parent.properties = {...newObj};
        };
        watch(() => {
            return props.value
        }, (newVal) => {
            const val = cloneDeep(newVal);
            data.value = addExtraInfo(val);
        }, {
            immediate: true,
            deep: true
        });
        watch(() => {
            return data.value
        }, (newVal) => {
            const newObj = removeExtraInfo(cloneDeep(newVal));
            console.log('832 emit change', newObj);
            emit('change', newObj);
        }, {
            immediate: true,
            deep: true
        });
        const renderAction = (isRoot: any, isFirst: any, isLast: boolean, keyIndex: number, parent: any) => {
            return <div class={'action'}>
                <Actions
                    isRoot={isRoot}
                    isFirst={isFirst || false}
                    isLast={isLast || false}
                    onMoveDown={moveDown.bind(this, keyIndex, parent)}
                    onMoveUp={moveUp.bind(this, keyIndex, parent)}
                    onCopy={copy.bind(this, keyIndex, parent)}/>
            </div>
        }
        const renderExtraAction = (isRoot: any, keyIndex: number, parent: any, tree: any) => {
            return <div class={'extraAction'}>
                <ExtraActions
                    isRoot={isRoot}
                    value={tree}
                    onAddDesc={addDesc.bind(this, tree)}
                    onDel={del.bind(this, keyIndex, parent)}
                    onSetRequire={setRequire.bind(this, keyIndex, parent)}/>
            </div>
        }
        const renderDataTypeSetting = (tree) => {
            return <>
                <DataTypeSetting refsOptions={props.refsOptions}
                                 value={tree}
                                 onChange={dataTypeChange.bind(this, tree)}/>
                {isObject(tree.type) ?
                    <span class={'baseInfoSpace'}>{`{${Object.keys(tree.properties || {}).length}}`}</span> : null}
            </>
        }
        const renderKeyName = (depth: number, keyName: string, keyIndex: number, parent: any) => {
            if (depth === 1) return null;
            return <>
                <span class={'baseInfoKey'}
                      contenteditable={true}
                      onInput={updateKeyName.bind(this, keyName, keyIndex, parent)}>
                    {keyName}
                </span>
                <span class={'baseInfoSpace'}>:</span>
            </>
        }
        const renderExpandIcon = (isExpand: boolean, tree: any) => {
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
                <div class={'leafNodeHorizontalLine'}
                     style={{left: `${(options.depth - 1) * treeLevelWidth + 8}px`}}/>
                <div class={'baseInfo'}>
                    {renderKeyName(options.depth, options.key, options.index, options.tree)}
                    {renderDataTypeSetting(options.value)}
                </div>
                {renderAction(options.isRoot, options.isFirst, options.isLast, options.index, options.tree)}
                {renderExtraAction(options.isRoot, options.index, options.tree, options.value)}
            </div>)
        }
        const renderDirectoryText = (options) => {
            const {depth, tree, keyName, keyIndex, isFirst, isLast, parent, isRoot, isExpand} = options;
            return <div class={'directoryText'}
                        style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                {renderHorizontalLine(depth)}
                <div class={'baseInfo'}>
                    {renderExpandIcon(isExpand, tree)}
                    {renderKeyName(depth, keyName, keyIndex, parent)}
                    {renderDataTypeSetting(tree)}
                    {isObject(tree.type) ? <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/> : null}
                </div>
                {renderAction(isRoot, isFirst, isLast, keyIndex, parent)}
                {renderExtraAction(isRoot, keyIndex, parent, tree)}
            </div>
        }
        const renderTree = (tree: any, option: any) => {
            if (!tree) {
                return null
            }
            const {keyIndex, parent, isFirst, isLast, keyName} = option;
            // 渲染对象类型节点
            if (isObject(tree.type)) {
                const {isRoot, isExpand, depth} = tree?.extraViewInfo || {};
                return <div class={{'directoryNode': true, "rootNode": isRoot}}>
                    {renderDirectoryText({
                        depth,
                        tree,
                        keyName,
                        keyIndex,
                        isFirst,
                        isLast,
                        parent,
                        isRoot,
                        isExpand
                    })}
                    <div class={{
                        'directoryContainer': tree,
                        'directoryContainerExpand': isExpand,
                        'directoryContainerFold': !isExpand
                    }}>
                        {
                            tree.properties && Object.entries(tree.properties).map(([key, value]: any, index: number, arr: any) => {
                                const isFirst = index === 0;
                                const isLast = index === arr.length - 1;
                                const depth = value.extraViewInfo.depth;
                                if (isObject(value.type)) {
                                    return renderTree(value, {
                                        keyName: key,
                                        keyIndex: index,
                                        tree: true,
                                        parent: tree,
                                        isFirst: isFirst,
                                        isLast: isLast,
                                    });
                                } else {
                                    return renderNormalType({
                                        index: key,
                                        depth: depth,
                                        key: key,
                                        tree: tree,
                                        value: value,
                                        isFirst: isFirst,
                                        isLast: isLast,
                                    })
                                }
                            })
                        }
                    </div>
                    <div class={'verticalLine'} style={{left: `${depth * treeLevelWidth + 8}px`}}></div>
                </div>
            }
            // 渲染数组类型节点
            // if (isArray(tree.type)) {
            //     const {isRoot, isExpand, depth} = tree?.extraViewInfo || {};
            //     return <div class={{'directoryNode': true, "rootNode": isRoot}}>
            //         {renderDirectoryText({depth, tree, keyName, keyIndex, isFirst, isLast, parent, isRoot, isExpand})}
            //         <div class={{
            //             'directoryContainer': tree,
            //             'directoryContainerExpand': isExpand,
            //             'directoryContainerFold': !isExpand
            //         }}>
            //             {
            //                 tree.items && Object.entries(tree.items.properties).map(([key, value]: any, index: number, arr: any) => {
            //                     const isFirst = index === 0;
            //                     const isLast = index === arr.length - 1;
            //                     const depth = value.extraViewInfo.depth;
            //                     if (isObject(value.type)) {
            //                         return renderTree(value, {
            //                             keyName: key,
            //                             keyIndex: index,
            //                             tree: true,
            //                             parent: tree,
            //                             isFirst: isFirst,
            //                             isLast: isLast,
            //                         });
            //                     } else {
            //                         return renderNormalType({
            //                             index: key,
            //                             depth: depth,
            //                             key: key,
            //                             tree: tree,
            //                             value: value,
            //                             isFirst: isFirst,
            //                             isLast: isLast,
            //                             isRoot: isRoot,
            //                         })
            //                     }
            //                 })
            //             }
            //         </div>
            //         <div class={'verticalLine'} style={{left: `${depth * treeLevelWidth + 8}px`}}></div>
            //     </div>
            // }
        }
        return () => (
            <div class={'schemaEditor-content'} style={props.contentStyle}>
                {renderTree(data.value, {})}
            </div>
        )
    }
})
