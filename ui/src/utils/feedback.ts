/**
 * 接入乐研评论反馈系统
 * @description 文档可参考 http://leyan-dev.rysaas.cn/docs/feedback/
 * */
import {initState, render, onHrefChange, helpChange, unmount} from '@leyan/feedback';
import '@leyan/feedback/lib/style.css'; // 默认样式

const renderfeedback = (currentUser:any) => {
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
                name: '模块1',
                desc: '模块描述',
                url: '模块url', // 点击模块跳转页面使用
                arrowHide: true, // 是否显示先后顺序关系
                list: [
                    {
                        label: '功能描述1',
                        icon: '连接地址',
                        popover: {
                            content: '单行文本',
                        },
                    },
                ],
            },
            {
                name: '模块2',
                desc: '模块描述',
                url: '模块url', // 点击模块跳转页面使用
                arrowHide: true, // 是否显示先后顺序关系
                list: [
                    {
                        label: '功能描述1',
                        icon: '连接地址',
                        popover: {
                            content: ['多行文本', '多行文本'],
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

        // 初始化数据
        initState({
            token: '32253824181731328', // 由乐研生成的应用授权Token
            api: process.env.NODE_ENV === 'production' ? 'https://leyan.nancalcloud.com/api/v1' : 'http://leyan-test.rysaas.cn/api/v1', // 请求地址可自行代理,乐研生产API: https://leyan.nancalcloud.com/api/v1
            help: null,
            userInfo: userInfo,
            zIndex: 1000,
            // filterHelp: filterHelp, // 自定义过滤显示帮助指引
            // onHelpChange: onHelpChange, // 帮助指引切换钩子
            // buildModule: buildModule, // 构建模块名称
        });

        // 渲染系统
        render();

        // 监听路由链接变化
        // onHrefChange((e) => {
        //     if (e.to == '/login') {
        //         unmount();
        //     } else {
        //         helpChange('key'); // 获取优先级：key > url >  tab_index
        //     }
        // });

    } catch (e) {
        console.error('乐研评论反馈系统，渲染报错：', e)
    }

};


export default renderfeedback;
