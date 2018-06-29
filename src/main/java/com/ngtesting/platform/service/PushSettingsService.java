package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;

import java.util.Map;

public interface PushSettingsService extends BaseService {
    void pushUserSettings(TstUser TstUser);

    void pushMyOrgs(TstUser TstUser);

    void pushOrgSettings(TstUser TstUser);

    void pushRecentProjects(TstUser TstUser);

    void pushPrjSettings(TstUser TstUser);

    void sendMsg(Integer userId, Map ret);
}
