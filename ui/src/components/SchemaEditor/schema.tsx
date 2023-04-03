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
import {cloneByJSON} from "@/utils/object";
import {removeExtraInfo, addExtraInfo, isArray, isObject} from './utils'

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
        const dataTypeChange = (tree:any, newVal) => {
            const [key, value]:any = Object.entries(newVal)[0];
            if (key) {
                tree.type = key;
                tree = Object.assign(tree, value);
            }
            console.log('832', key, value,tree)
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
        const addDesc = (keyIndex: any, parent: any) => {
            const keys = Object.keys(parent.properties);
            const key = keys[keyIndex];
            // ::::todo 添加描述逻辑
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
            const val = cloneByJSON(newVal);
            data.value = addExtraInfo(val);
        }, {
            immediate: true,
            deep: true
        });

        watch(() => {
            return data.value
        }, (newVal) => {
            const newObj = removeExtraInfo(cloneByJSON(newVal));
            console.log('832 emit change',newObj);
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
        const renderExtraAction = (isRoot: any, keyIndex: number, parent: any) => {
            return <div class={'extraAction'}>
                <ExtraActions
                    isRoot={isRoot}
                    onAddDesc={addDesc.bind(this, keyIndex, parent)}
                    onDel={del.bind(this, keyIndex, parent)}
                    onSetRequire={setRequire.bind(this, keyIndex, parent)}/>
            </div>
        }

        const renderDataTypeSetting = (tree) => {
            return <DataTypeSetting refsOptions={props.refsOptions}
                                    value={tree}
                                    onChange={dataTypeChange.bind(this, tree)}/>
        }

        const treeLevelWidth = 24;
        const renderTree = (tree: any, option: any) => {
            if (!tree) {
                return null
            }
            const {keyIndex, parent, isFirst, isLast, keyName} = option;
            if (isObject(tree.type)) {
                const {isRoot, isExpand, depth} = tree?.extraViewInfo || {};
                return <div class={{'directoryNode': true, "rootNode": isRoot}}>
                    <div class={'directoryText'}
                         style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                        {!isRoot ?
                            <div class={'horizontalLine'}
                                 style={{left: `${(depth - 1) * treeLevelWidth + 8}px`}}></div> : null}

                        <div class={'baseInfo'}>
                            {isExpand ? <DownOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/> : null}
                            {!isExpand ?
                                <RightOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/> : null}
                            {!isRoot ? <span class={'baseInfoKey'}
                                             contenteditable={true}
                                             onInput={updateKeyName.bind(this, keyName, keyIndex, parent)}>{keyName}</span> : null}
                            {!isRoot ? <span class={'baseInfoSpace'}>:</span> : null}
                            {renderDataTypeSetting(tree)}
                            <span class={'baseInfoSpace'}>{`{${Object.keys(tree.properties || {}).length}}`}</span>
                            <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/>
                        </div>
                        {renderAction(isRoot, isFirst, isLast, keyIndex, parent)}
                        {renderExtraAction(isRoot, keyIndex, parent)}
                    </div>
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
                                    return (<div
                                        key={index}
                                        class={'leafNode'} style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                                        <div class={'leafNodeHorizontalLine'}
                                             style={{left: `${(depth - 1) * treeLevelWidth + 8}px`}}/>
                                        <div class={'baseInfo'}>
                                            <span class={'baseInfoKey'} contenteditable={true}
                                                  onInput={updateKeyName.bind(this, key, index, tree)}>{key}</span>
                                            <span class={'baseInfoSpace'}>:</span>
                                            {renderDataTypeSetting(value)}
                                        </div>
                                        {renderAction(isRoot, isFirst, isLast, index, tree)}
                                        {renderExtraAction(isRoot, index, tree)}
                                    </div>)
                                }
                            })
                        }
                    </div>
                    <div class={'verticalLine'} style={{left: `${depth * treeLevelWidth + 8}px`}}></div>
                </div>
            }
            if (isArray(tree.type)) {
                const {isRoot, isExpand, depth} = tree?.extraViewInfo || {};
                return <div class={{'directoryNode': true, "rootNode": isRoot}}>
                    <div class={'directoryText'}
                         style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                        {!isRoot ? <div class={'horizontalLine'}
                                        style={{left: `${(depth - 1) * treeLevelWidth + 8}px`}}></div> : null}
                        <div class={'baseInfo'}>
                            {isExpand ?
                                <DownOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/> : null}
                            {!isExpand ?
                                <RightOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/> : null}
                            {!isRoot ? <span class={'baseInfoKey'}
                                             contenteditable={true}
                                             onInput={updateKeyName.bind(this, keyName, keyIndex, parent)}>{keyName}</span> : null}
                            {!isRoot ? <span class={'baseInfoSpace'}>:</span> : null}
                            {renderDataTypeSetting(tree)}
                            <span class={'baseInfoSpace'}>{`{${Object.keys(tree.properties || {}).length}}`}</span>
                            <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/>
                        </div>

                        {renderAction(isRoot, isFirst, isLast, keyIndex, parent)}
                        {renderExtraAction(isRoot, keyIndex, parent)}

                    </div>
                    <div class={{
                        'directoryContainer': tree,
                        'directoryContainerExpand': isExpand,
                        'directoryContainerFold': !isExpand
                    }}>
                        {
                            tree.items && Object.entries(tree.items.properties).map(([key, value]: any, index: number, arr: any) => {
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
                                    return (<div
                                        key={index}
                                        class={'leafNode'} style={{'paddingLeft': `${depth * treeLevelWidth}px`}}>
                                        <div class={'leafNodeHorizontalLine'}
                                             style={{left: `${(depth - 1) * treeLevelWidth + 8}px`}}/>
                                        <div class={'baseInfo'}>
                                            <span class={'baseInfoKey'} contenteditable={true}
                                                  onInput={updateKeyName.bind(this, key, index, tree)}>{key}</span>
                                            <span class={'baseInfoSpace'}>:</span>
                                            {renderDataTypeSetting(value)}
                                        </div>
                                        {renderAction(isRoot, isFirst, isLast, index, tree)}
                                        {renderExtraAction(isRoot, index, tree)}
                                    </div>)
                                }
                            })
                        }
                    </div>
                    <div class={'verticalLine'} style={{left: `${depth * treeLevelWidth + 8}px`}}></div>
                </div>
            }
        }
        return () => (
            <div class={'schemaEditor-content'} style={props.contentStyle}>
                {renderTree(data.value, {})}
            </div>
        )
    }
})
