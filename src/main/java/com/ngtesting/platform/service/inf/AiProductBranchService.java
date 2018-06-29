package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.AiProductBranch;

import java.util.List;

public interface AiProductBranchService extends BaseService {

	List<AiProductBranch> listForProductBranchVo(Long projectId);

}
