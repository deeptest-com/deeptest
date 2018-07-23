package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface UserService {
    List<TstUser> list(Integer orgId, String keywords, String disabled, int pageNum, int pageSize);

    TstUser get(Integer id);

    TstUser getByToken(String token);
    TstUser getByPhone(String token);
    TstUser getByEmail(String email);

    TstUser invitePers(TstUser user, TstUser vo, List<TstOrgGroupUserRelation> relations);

//    TstUser save(TstUser vo, Integer orgId);

    void update(TstUser record);

    void setDefaultOrg(TstUser user, Integer orgId);

    void setDefaultPrj(TstUser user, Integer prjId);

    List<TstUser> search(Integer orgId, String keywords, String exceptIds);

}
