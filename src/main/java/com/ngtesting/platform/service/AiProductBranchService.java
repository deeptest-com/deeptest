package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiProductBranch;
import com.ngtesting.platform.vo.AiProductBranchVo;

import java.util.List;

public interface AiProductBranchService extends BaseService {

	List<AiProductBranchVo> listForProductBranchVo(Long projectId);

	List<AiProductBranchVo> genVos(List<AiProductBranch> pos);

	AiProductBranchVo genVo(AiProductBranch po);
}
