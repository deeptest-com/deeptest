import {
    defineComponent,
    ref,
    onMounted,
    onUnmounted,
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
import SettingPropsModal from './SettingPropsModal.vue';
import {computePosition} from '@floating-ui/dom';
import {cloneByJSON} from "@/utils/object";

function isLeafNode(type: string) {
    return ['string', 'boolean', 'integer', 'number'].includes(type)
}

function isObject(type: string) {
    return type === 'object';
}

export default defineComponent({
    name: 'SchemeEditor',
    props: {
        value: Object,
        contentStyle: Object
    },
    setup(props) {
        const data: any = ref(null);
        const expandIt = (tree: any, e: any) => {
            if (tree?.extraViewInfo) {
                tree.extraViewInfo.isExpand = !tree.extraViewInfo.isExpand;
            }
        }
        const visible = ref(false);

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
            data.value = adaptValue(data.value);
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

        function handleModalOk() {
            visible.value = false;
        }

        function handleModalCancel() {
            visible.value = false;
        }

        const floatingCon: any = ref(null);
        const floating: any = ref(null);
        const floatingArrow: any = ref(null);
        const activeTree = ref(null);
        const showSettingPropsModal = (tree: any, e: any) => {
            visible.value = true;
            activeTree.value = tree;
            computePosition(e.target, floating.value, {
                placement: 'right-start',
                middleware: [],
            }).then(({x, y, middlewareData}) => {
                Object.assign(floating.value.style, {
                    left: `${8 + x}px`,
                    top: `${y}px`,
                });

            });
        };


        onMounted(() => {
            // 添加点击事件监听器
            document.addEventListener('click', (event) => {
                // 如果单击事件不是发生在目标元素或其后代元素上
                // visible.value = floatingCon?.value.contains(event.target);
                const target: any = event?.target;
                if (target?.className && target?.className?.includes && target?.className?.includes('setDataTypeAction')) {
                    return;
                }
                if (!floatingCon?.value?.contains(event.target)) {
                    visible.value = false;
                }
            });
        })

        onUnmounted(() => {
            console.log('销毁')
        })

        // 适配数据结构
        function adaptValue(val) {
            if (!val) {
                return null
            }
            val.extraViewInfo = {
                "isExpand": true,
                "isRoot": true,
                "name": "root",
                "depth": 1,
            };

            function fn(obj: any, depth) {
                if (obj.properties && obj.type === 'object') {
                    Object.entries(obj.properties).forEach(([key, value]: any) => {
                        value.extraViewInfo = {
                            "isExpand": true,
                            "isRoot": false,
                            "name": key,
                            "depth": depth,
                        }
                        if (value.type === 'object') {
                            fn(value, depth + 1);
                        }
                    })
                }
            }

            fn(val, 2);
            return val;
        }

        watch(() => {
            return props.value
        }, (newVal) => {
            const val = cloneByJSON(newVal);
            data.value = adaptValue(val);
            console.log('watch props value 832', data.value);
        }, {
            immediate: true,
            deep: true
        });

        watch(() => {
            return data.value
        }, (newVal) => {
            console.log('watch data value 832', newVal);
        }, {
            immediate: true,
            deep: true
        });

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
                            {isExpand ?
                                <DownOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/> : null}
                            {!isExpand ?
                                <RightOutlined onClick={expandIt.bind(this, tree)} class={'expandIcon'}/> : null}
                            {!isRoot ? <span class={'baseInfoKey'}
                                             contenteditable={true}
                                             onInput={updateKeyName.bind(this, keyName, keyIndex, parent)}>{keyName}</span> : null}
                            {!isRoot ? <span class={'baseInfoSpace'}>:</span> : null}
                            <a href="javascript:void(0)"
                               onClick={showSettingPropsModal.bind(this, tree)}
                               class={[tree.type, 'setDataTypeAction']}
                            >{tree.type}</a>
                            <span class={'baseInfoSpace'}>{`{${Object.keys(tree.properties || {}).length}}`}</span>
                            <PlusOutlined onClick={addProps.bind(this, tree)} class={'addIcon'}/>
                        </div>
                        <div class={'action'}>
                            <Actions
                                isRoot={isRoot}
                                isFirst={isFirst}
                                isLast={isLast}
                                onMoveDown={moveDown.bind(this, keyIndex, parent)}
                                onMoveUp={moveUp.bind(this, keyIndex, parent)}
                                onCopy={copy.bind(this, keyIndex, parent)}/>
                        </div>
                        <div class={'extraAction'}>
                            <ExtraActions
                                isRoot={isRoot}
                                onAddDesc={addDesc.bind(this, keyIndex, parent)}
                                onDel={del.bind(this, keyIndex, parent)}
                                onSetRequire={setRequire.bind(this, keyIndex, parent)}/>
                        </div>
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
                                            <a class={[value.type, 'setDataTypeAction']}
                                               onClick={showSettingPropsModal.bind(this, value)}
                                               href='javascript:void(0)'>{value.type}</a>
                                        </div>
                                        <div class={'action'}>
                                            <Actions
                                                isFirst={isFirst}
                                                isLast={isLast}
                                                isRoot={false}
                                                onMoveDown={moveDown.bind(this, index, tree)}
                                                onCopy={copy.bind(this, index, tree)}
                                                onMoveUp={moveUp.bind(this, index, tree)}/>
                                        </div>
                                        <div class={'extraAction'}>
                                            <ExtraActions
                                                isRoot={false}
                                                onAddDesc={addDesc.bind(this, index, tree)}
                                                onDel={del.bind(this, index, tree)}
                                                onSetRequire={setRequire.bind(this, index, tree)}/>
                                        </div>
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
                <div ref={floatingCon}>
                    <div
                        class={'floatingSetting'}
                        ref={floating}
                        style={{
                            position: 'absolute',
                            display: visible.value ? 'block' : 'none',
                        }}
                    >
                        <a-card
                            bodyStyle={{padding: '0 16px 16px 16px'}}
                            class={'floatingSetting-card'}
                            title={null}>
                            <SettingPropsModal
                                onOk={handleModalOk}
                                onCancel={handleModalCancel}
                                value={activeTree.value}
                                visible={visible.value}/>
                        </a-card>
                        <div ref={floatingArrow} class="floatingSetting-arrow"></div>
                    </div>
                </div>
            </div>
        )
    }
})
