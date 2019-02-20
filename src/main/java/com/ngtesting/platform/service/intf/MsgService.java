package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.utils.MsgUtil;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface MsgService extends BaseService {
    List<TstMsg> list(Integer userId, Boolean isRead, String keywords);

    void createForTask(TstUser optUser, TstTask task, MsgUtil.HistoryMsgTemplate template, Object... params);

    @Transactional
        // 针对经办人和监听者
    void createForIssue(TstUser optUser, IsuIssue issue, MsgUtil.HistoryMsgTemplate template, Object... params);

    Boolean delete(Integer msgId, Integer userId);

    Boolean markRead(Integer id, Integer userId);
    void markAllRead(Integer id);

}
