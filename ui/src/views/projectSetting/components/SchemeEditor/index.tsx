import {computed, defineComponent, defineEmits, defineProps, reactive, Ref, ref, UnwrapRef, onMounted} from 'vue';
import './index.less';

function isLeafNode(type) {
    return ['string', 'boolean', 'integer', 'number'].includes(type)
}

function isObject(type) {
    return type === 'object';
}

export default defineComponent({
    name: 'SchemeEditor',
    setup() {
        const data = reactive({
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                // "address": {
                //     "$ref": "#/components/schemas/Address"
                // },
                "age": {
                    "type": "integer",
                    "format": "int32",
                    "minimum": 0
                },
                'obj1': {
                    "type": "object",
                    "required": [
                        "name"
                    ],
                    "properties": {
                        "name1": {
                            "type": "string"
                        },
                        "age1": {
                            "type": "integer",
                            "format": "int32",
                            "minimum": 0
                        }
                    }
                }
            }
        });

        // onMounted(() => {
        //
        // });

        const renderTree = (tree) => {
            console.log(832832222,tree.type);
            if (!isObject(tree.type)) {
                console.log('')
                return (<div class={'li'}>{tree.type}</div>)
            }
            return <div class={'ul'}>
                {
                    Object.entries(tree.properties).map(([key, value]) => {
                        console.log(832,value.type);
                        renderTree(value);
                    })
                }
            </div>
        }


        return () => (
            <div>
                {renderTree(data)}
            </div>
        )
    }

})
