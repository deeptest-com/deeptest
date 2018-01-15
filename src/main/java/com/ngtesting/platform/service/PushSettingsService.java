package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.UserVo;

import java.util.Map;

public interface PushSettingsService extends BaseService {
    void pushUserSettings(UserVo userVo);

    void pushMyOrgs(UserVo userVo);

    void pushOrgSettings(UserVo userVo);

    void pushRecentProjects(UserVo userVo);

    void pushPrjSettings(UserVo userVo);

    void sendMsg(Long userId, Map ret);
}
