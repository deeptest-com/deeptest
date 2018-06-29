package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.inf.AiAudioTypeService;
import com.ngtesting.platform.model.AiAudioType;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AiAudioTypeServiceImpl extends BaseServiceImpl implements AiAudioTypeService {

    @Override
    public List<AiAudioType> listAudioTypeVo(Long projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(AiAudioType.class);
//
////        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("displayOrder"));
//
//        List<AiAudioType> ls = findAllByCriteria(dc);
//        List<AiAudioType> vos = genVos(ls);

        return null;
    }
}

