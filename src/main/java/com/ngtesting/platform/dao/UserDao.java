package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface UserDao {
    List<TstUser> query(@Param("orgId") Integer orgId,
                        @Param("keywords") String keywords,
                        @Param("disabled") Boolean disabled);

    List<TstUser> search(@Param("orgId") Integer orgId,
                         @Param("keywords") String keywords,
                         @Param("exceptIds") String exceptIds);

    TstUser get(@Param("id") Integer id);
    TstUser getByEmail(@Param("email") String email);
    TstUser getByEmailAndPassword(@Param("email") String email,
                                  @Param("password") String password);
    TstUser getByToken(@Param("token") String token);
    TstUser getByPhone(@Param("phone") String phone);

    String getSalt(@Param("id") Integer id);

    void save(TstUser record);
    void update(TstUser record);

    void modifyProp(@Param("id") Integer id,
                    @Param("prop") String prop,
                    @Param("value") String value);

    void setDefaultOrg(@Param("id") Integer id,
                       @Param("orgId") Integer orgId,
                       @Param("orgName") String orgName);

    void setDefaultPrj(@Param("id") Integer id,
                       @Param("prjId") Integer prjId,
                       @Param("prjName") String prjName);

    void setLeftSize(TstUser user);

    List<TstUser> getProjectUsers(@Param("prjId") Integer prjId, @Param("numb") Integer numb);

    void setDefaultOrgPrjToNullForDelete(@Param("orgId") Integer orgId);
}
