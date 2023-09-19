/**
 * 接入乐研评论反馈系统
 * @description 文档可参考 http://leyan-dev.rysaas.cn/docs/feedback/
 * */
import {initState, render, onHrefChange, helpChange, unmount} from '@leyan/feedback';
import '@leyan/feedback/lib/style.css'; // 默认样式
// import PutAway from '@/assets/images/put-away.png';
const renderfeedback = (currentUser: any) => {
    try {
        /** 用户信息 */
        const userInfo = {
            userName: currentUser?.value?.username,
            userKey: currentUser?.value?.username,
            email: currentUser?.value?.email,
        };
        /** 指引配置 */
        const help = [
            {
                name: '后端开发',
                desc: '后端开发可以在平台设计和定义遵循OpenAPI 3.0规范的接口，进行接口自测、发布接口文档、debug调试等工作。',
                arrowHide: false,
                list: [
                    {
                        label: '接口定义',
                        icon: require('@/assets/feedback/interface.svg'),
                        popover: {
                            content: [
                                {
                                    title: '在功能开发前，前后端先约定接口协议，便于前后端分离并行开发，提高开发效率。',
                                    content: [
                                        '接口定义遵循OpenAPI 3.0规范，可基于endpoint路径定义多种请求方法，定义不同请求方法的请求参数、响应数据结构等。' ,
                                        '支持基本数据类型、自定义复合类型，可引用数据组件。' ,
                                        '支持从swagger等常见格式接口数据批量导入接口定义',
                                    ]
                                }
                            ],
                        },
                    },
                    {
                        label: '发布接口文档',
                        icon: require('@/assets/feedback/doc.svg'),
                        popover: {
                            content: [
                                {
                                    title: '文档功能可以为接口调用方提供可读性更高的接口文档，可发布不同版本的接口文档，并支持分享链接',
                                    content: [
                                        '通过 接口定义 -> 文档 功能可预览单接口文档。',
                                        '通过“接口文档”功能可查看最新的全量接口文档及所有历史版本。',
                                        '可以根据接口文档自动生成请求代码、数据结构体代码等',
                                    ]
                                }
                            ],
                        },
                    },
                    {
                        label: '开发调试',
                        icon: require('@/assets/feedback/code.svg'),
                        popover: {
                            content: [
                                '开发过程中需要对依赖的第三方接口进行调用调试。在没有定义接口的情况下，可以通过“快捷调试”调用第三方接口'
                            ],
                        },
                    },
                    {
                        label: '接口自测',
                        icon: require('@/assets/feedback/t-c.svg'),
                        popover: {
                            content: [
                                {
                                    title: '开发完成的接口需要先充分自测再交给测试。',
                                    content: [
                                        '通过 接口定义 -> 调试 功能调用自己开发的接口进行验证',
                                        '通过 接口定义 -> 用例 功能将常见请求参数及期望的响应结果沉淀为单接口用例，便于反复验证。'
                                    ]
                                }
                            ],
                        },
                    },
                    {
                        label: '快捷调试',
                        icon: require('@/assets/feedback/debug.svg'),
                        popover: {
                            content: '测试过程中发现的bug、线上问题需尽快修复，可以通过“快捷调试”功能排查定位后端接口问题，提高bug修复效率',
                        },
                    },
                ],
            },
            {
                name: '前端开发',
                desc: '前端开发可以在平台查看接口文档，通过mock功能和后端并行开发，开发过程中做接口调试、前后端接口联调，并通过快捷调试进行bug的快速定位和修复。',
                list: [

                    {
                        label: '查看接口文档',
                        icon: require('@/assets/feedback/doc.svg'),
                        popover: {
                            content: [
                                {
                                    title: '在后端接口开发完成前，通过查看接口文档，尽快展开前端的设计和开发工作。',
                                    content: [
                                        '通过 接口定义 -> 文档 功能查看单接口文档。',
                                        '通过“接口文档”入口可查看最新的全量接口文档以及所有历史版本。',
                                        '可以根据接口文档自动生成请求代码、数据结构体代码等',
                                    ],
                                },
                            ],
                        },
                    },

                    {
                        label: '接口Mock',
                        icon: require('@/assets/feedback/mock.svg'),
                        popover: {
                            content: [
                                {
                                    title: '在后端接口开发完成前，通过Mock功能模拟后端接口返回，提前开发前端功能，提高开发效率',
                                    content: [
                                        '调用接口Mock地址，系统将根据Mock规则智能生成Mock响应数据',
                                        '通过高级Mock功能，精准指定特定请求参数下期望的返回数据',
                                    ],
                                },
                            ],
                        },
                    },
                    {
                        label: '开发联调',
                        icon: require('@/assets/feedback/code.svg'),
                        popover: {
                            content: [
                                '前后端各自开发完成后需要进行联调，可以通过“接口定义 -> 调试” 或 “快捷调试” 功能调用验证后端接口。'
                            ],
                        },
                    },

                    {
                        label: '快捷调试',
                        icon: require('@/assets/feedback/debug.svg'),
                        popover: {
                            content: [
                                '测试过程中发现的bug、线上问题需尽快修复，可以通过“快捷调试”功能排查定位后端接口问题，提高bug修复效率'
                            ],
                        },
                    },

                ],
            },
            {
                name: '测试',
                desc: '测试过程遵循金字塔模式，先做单接口测试、再针对特定功能做场景测试，针对系统的整体测试以及例行功能回归验证。平台通过接口用例、场景测试、测试计划等功能为测试同学提供支持。',
                list: [
                    {
                        label: '查看接口文档',
                        icon: require('@/assets/feedback/doc.svg'),
                        popover: {
                            content: [
                                {
                                    title: '在后端接口开发完成前，通过查看接口文档，提前展开测试准备、测试用例编写等工作。',
                                    content: [
                                        '通过 接口定义 -> 文档 功能查看单接口文档。',
                                        '通过“接口文档”入口可查看最新的全量接口文档以及所有发布过的历史版本。',
                                        '可以根据接口文档自动生成请求代码、数据结构体代码等',
                                    ],
                                },
                            ],
                        },
                    },

                    {
                        label: '接口测试',
                        icon: require('@/assets/feedback/t-c.svg'),
                        popover: {
                            content: [
                                '针对接口定义，基于不同的接口请求参数及预期的响应结果，可以创建多个接口用例，用于单接口测试'
                            ],
                        },
                    },

                    {
                        label: '场景测试',
                        icon: require('@/assets/feedback/s-t.svg'),
                        popover: {
                            content: [
                                '针对特定功能和场景，将多个接口请求进行编排，进行接口自动化场景测试。',
                                '通过“测试开发”功能进行场景测试，场景用例中支持多种常见处理器，如条件分支、循环、断言等。'
                            ],
                        },
                    },

                    {
                        label: '测试计划',
                        icon: require('@/assets/feedback/plan.svg'),
                        popover: {
                            content: [
                                '针对一次项目发布、产品版本发布或例行系统回归，可创建完整的测试计划。一个测试计划中可以包含多个场景用例，批量执行，并生成统一的测试报告。'
                            ],
                        },
                    },

                    {
                        label: '测试报告',
                        icon: require('@/assets/feedback/note.svg'),
                        popover: {
                            content: [
                                '测试计划的每次执行都会生成一份测试报告，场景用例执行的结果也可以保存为测试报告。',
                                '“测试报告”功能提供了统一查看、管理所有测试结果报告的统一入口'
                            ],
                        },
                    },
                ],
            },
            {
                name: '项目管理员',
                desc: '管理员或开发/测试负责人，通过“项目设置”功能入口进行项目成员维护、环境管理、服务管理，设置Mock策略、通过数据池共享数据等。',
                list: [
                    {
                        label: '环境管理',
                        icon: require('@/assets/feedback/e-m.svg'),
                        popover: {
                            content: [
                                {
                                    title: '新功能开发测试需要通过不同环境下的逐层验证。通常包括开发环境、测试环境、UAT环境、生产环境。',
                                    content: [
                                        '各个环境中可添加服务、设置服务前置URL，设置通用的环境变量',
                                        '设置各环境通用的全局变量/参数',
                                        '在接口请求调用时，通过切换环境便捷改变请求前置URL。',
                                    ],
                                },
                            ],
                        },
                    },

                    {
                        label: '数据池',
                        icon: require('@/assets/feedback/db.svg'),
                        popover: {
                            content: [
                                '通过数据池可以维护项目中可共享的测试数据。在接口调试、接口用例、场景用例中可以便捷引用数据池中定义好的数据。'
                            ],
                        },
                    },

                    {
                        label: '服务管理',
                        icon: require('@/assets/feedback/s-m.svg'),
                        popover: {
                            content: [
                                {
                                    title: '接口定义归属于某个服务，每个环境中可以添加不同的服务。',
                                    content: [
                                        '服务下可以维护共享的数据组件，在接口中可以直接引用，避免重复定义相同的数据结构。',
                                        '服务下也可以定义统一的安全策略。',
                                    ],
                                },
                            ],
                        },
                    },
                    {
                        label: '项目成员',
                        icon: require('@/assets/feedback/user.svg'),
                        popover: {
                            content: [
                                '可以邀请/添加项目成员，维护项目成员角色等信息'
                            ],
                        },
                    },
                    {
                        label: 'Mock设置',
                        icon: require('@/assets/feedback/m-s.svg'),
                        popover: {
                            content: [
                                '可以自定义Mock响应数据生成的优先策略。Mock设置中可指定响应示例优先，在未匹配到高级mock规则的情况下，将直接返回接口定义中的响应示例数据，没有示例数据再使用智能Mock生成数据并返回。',
                            ],
                        },
                    },
                    {
                        label: '自动同步',
                        icon: require('@/assets/feedback/sync.svg'),
                        popover: {
                            content: [
                                '开启Swagger自动同步，系统将从指定的Swagger地址中定时自动同步接口定义到当前项目中的指定分类目录下。',
                            ],
                        },
                    },
                ],
            },

        ];
        /** 自定义过滤 */
        const filterHelp = (list, e) => {
            if (e.to !== '/login') {
                return list;
            } else {
                return [];
            }
        };
        /** 监听指引模块切换 */
        const onHelpChange = (e) => {
            window.open(e.url);
        };

        /** 构建模块名称 */
        const buildModule = ({url}) => {
            if (url.pathname.startsWith('/aaa')) {
                return '模块1';
            }
        };

        const isProd = window.location.href.includes('leyanapi.nancalcloud.com');
        // 初始化数据
        initState({
            token: '32253824181731328', // 由乐研生成的应用授权Token
            api: isProd ? 'https://leyan.nancalcloud.com/api/v1' : 'http://leyan-test.rysaas.cn/api/v1',
            help: help,
            userInfo: userInfo,
            zIndex: 1000,
            // filterHelp: filterHelp, // 自定义过滤显示帮助指引
            // onHelpChange: onHelpChange, // 帮助指引切换钩子
            // buildModule: buildModule, // 构建模块名称
        });

        // 渲染系统
        render();

        // 监听路由链接变化
        onHrefChange((e) => {
            if (e.to.includes('user/login')) {
                unmount();
            }
        });

    } catch (e) {
        console.error('乐研评论反馈系统，渲染报错：', e)
    }

};


export default renderfeedback;
