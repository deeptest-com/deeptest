package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.vo.IsuJqlColumn;

import java.util.List;

public interface IssueJqlColumnService extends BaseService {

    List<IsuJqlColumn> loadColumns(TstUser user);

    String buildDefaultColStr(TstUser user);
}
