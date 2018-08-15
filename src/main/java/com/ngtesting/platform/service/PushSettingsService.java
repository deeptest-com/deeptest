package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;

import javax.servlet.http.HttpServletRequest;
import java.util.Map;

public interface PushSettingsService extends BaseService {
    void pushUserSettings(TstUser TstuserUser);

    void pushMyOrgs(TstUser user);

    void pushOrgSettings(TstUser user, HttpServletRequest request);

    void pushRecentProjects(TstUser user);

    void pushPrjSettings(TstUser user);

    void sendMsg(TstUser user, Map ret);
}
