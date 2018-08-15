package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;

import javax.servlet.http.HttpServletRequest;
import java.util.List;

public interface UserService {
    List<TstUser> list(Integer orgId, String keywords, Boolean disabled, int pageNum, int pageSize);

    List<TstUser> listAllOrgUsers(Integer orgId);

    TstUser get(Integer id);

    TstUser getByToken(String token);
    TstUser getByPhone(String token);
    TstUser getByEmail(String email);

    TstUser invite(TstUser user, TstUser vo, List<TstOrgGroupUserRelation> relations);

    TstUser update(TstUser record);

    TstUser modifyProp(JSONObject json);

    void setDefaultOrg(TstUser user, Integer orgId);

    void setDefaultPrj(TstUser user, Integer prjId);

    List<TstUser> search(Integer orgId, String keywords, String exceptIds);

    TstUser setLeftSizePers(TstUser user, Integer left, String prop);

    List<TstUser> getProjectUsers(Integer orgId, Integer projectId);

}
