package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;

import java.util.List;

public interface IsuTqlService extends BaseService {
    List<TstVer> getFilters(String tql);

    Boolean save(Integer caseId, String name, String path, TstUser user);
    Boolean delete(Integer id, TstUser user);
}
