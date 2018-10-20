package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.vo.IsuJqlColumn;

import java.util.List;

public interface IsuJqlColumnService extends BaseService {

    List<IsuJqlColumn> loadColumns(TstUser user);

    String buildDefaultColStr(TstUser user);
}
