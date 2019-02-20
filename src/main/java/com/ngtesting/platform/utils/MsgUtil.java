package com.ngtesting.platform.utils;

import org.apache.commons.lang3.StringUtils;

public class MsgUtil extends StringUtils {

//    public enum MsgEntity {
//        test_plan("测试计划"),
//        test_task("测试任务"),
//        test_case("测试用例"),
//        issue("问题"),
//        issue_attachment("问题附件"),
//        issue_comments("问题备注");
//
//        MsgEntity(String msg) {
//            this.msg = msg;
//        }
//
//        public String msg;
//    }

    public enum MsgAction {
        create("创建"),
        rename("改名"),
        update("更新"),
        move("移动"),
        copy("复制"),
        delete("删除"),

        assign("分配经办人"),
        changeStatus("修改状态"),
        link("建立链接"),
        watch("关注问题"),
        unwatch("取消关注"),
        removeWatch("移除关注人"),
        changeWatch("修改关注列表"),

        attachment_upload("上传附件"),
        attachment_delete("删除附件"),

//        comments_add("新增注释"),
//        comments_update("修改注释"),
//        comments_delete("删除注释"),

        exe_result("标注执行结果");

        MsgAction(String msg) {
            this.msg = msg;
        }

        public String msg;
    }

    public enum HistoryMsgTemplate {
        opt_project("用户{0}{1}{3}{4}"),
        opt_entity("用户{0}{1}{2}"),
        exe_case("用户{0}标记执行状态为{1}"),

        create_task("用户{0}创建任务{1}"),
        update_task("用户{0}更新任务{1}"),

        start_task("任务{0}开始"),
        end_task("任务{0}结束"),

        create_issue("用户{0}创建问题{1}"),
        update_issue("用户{0}更新问题{1}"),
        update_issue_field("用户{0}更新问题{1}的字段{2}"),

        create_attament_for_issue("用户{0}为问题{1}上传附件{2}"),
        remove_attament_for_issue("用户{0}为问题{1}更新附件{2}"),
        create_comments_for_issue("用户{0}为问题{1}添加注释{2}"),
        update_comments_for_issue("用户{0}为问题{1}更新注释{2}"),
        remove_comments_for_issue("用户{0}为问题{1}移除注释{2}"),

        update_case_for_task("用户{0}更新任务{1}的用例");

        HistoryMsgTemplate(String msg) {
            this.msg = msg;
        }

        public String msg;
    }
    
}
