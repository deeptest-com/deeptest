package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.service.inf.AlertService;
import com.ngtesting.platform.service.inf.MsgService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class OptFacade {
    Log logger = LogFactory.getLog(OptFacade.class);

//    @Autowired
//    OptNews optNews;
//    @Autowired
//    OptQa optQa;
//    @Autowired
//    OptChat optChat;

    @Autowired
    MsgService msgService;
    @Autowired
    AlertService alertService;

    public void opt(JSONObject json, String userId) {
//        Map<String, Object> ret = new HashMap<>();
//        String type = json.getString("type");
//
//        try {
//            if (WsConstant.WS_OPEN.equals(type) || WsConstant.WS_TODO.equals(type)) {
//                List<TstMsg> msgs = msgService.list(Long.valueOf(userId), false);
//                List<TstAlert> alerts = alertService.list(Long.valueOf(userId), false);
//
//                ret.put("msgs", msgs);
//                ret.put("alerts", alerts);
//
//                ret.put("type", WsConstant.WS_MSG_AND_ALERT_LASTEST);
//            } else if (WsConstant.WS_OPT_ENTER_CHAT_ROOM.equals(type)) {
//                ret = optChat.enter(json);
//            } else if (WsConstant.WS_OPT_CHAT.equals(type)) {
//                ret = optChat.chat(json);
//            }
//        } catch (Exception e) {
//            logger.error(ExceptionUtil.GetExceptionInfo(e));
//            if (e.getCause() instanceof TimeoutException) {
//                ret.put("code", RespCode.BIZ_FAIL.getCode());
//                ret.put("msg", "操作超时！");
//            }else{
//                ret.put("msg", "操作出错，请稍后重试");
//            }
//        }
//
//        ret.put("code", 1);
//        if (ret.get("type") != null) {
//            scopeBean.sendMessageToClient(userId, new TextMessage(JSON.toJSONString(ret)));
//        }
    }

    public void opt(String act, String userId) {
        JSONObject json = new JSONObject();
        json.put("type", act);
        opt(json, userId);
    }

}
