package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface UserService {
    List<TstUser> list(Integer orgId, String keywords, Boolean disabled, int pageNum, int pageSize);

    List<TstUser> listAllOrgUsers(Integer orgId);

    TstUser get(Integer id);

    TstUser getByToken(String token);
    TstUser getByPhone(String token);
    TstUser getByEmail(String email);

    TstUser invitePers(TstUser user, TstUser vo, List<TstOrgGroupUserRelation> relations);

    void update(TstUser record);

    void setDefaultOrg(TstUser user, Integer orgId);

    void setDefaultPrj(TstUser user, Integer prjId);

    List<TstUser> search(Integer orgId, String keywords, String exceptIds);

    TstUser setLeftSizePers(TstUser user, Integer left, String prop);

    List<TstUser> getProjectUsers(Integer orgId, Integer projectId);
}
