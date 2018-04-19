package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.AiAudioType;
import com.ngtesting.platform.service.AiAudioTypeService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.AiAudioTypeVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiAudioTypeServiceImpl extends BaseServiceImpl implements AiAudioTypeService {

    @Override
    public List<AiAudioTypeVo> listAudioTypeVo(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiAudioType.class);

//        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("displayOrder"));

        List<AiAudioType> ls = findAllByCriteria(dc);
        List<AiAudioTypeVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public List<AiAudioTypeVo> genVos(List<AiAudioType> pos) {
        List<AiAudioTypeVo> vos = new LinkedList<>();

        for (AiAudioType langModel : pos) {
            AiAudioTypeVo vo = genVo(langModel);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiAudioTypeVo genVo(AiAudioType po) {
        AiAudioTypeVo vo = new AiAudioTypeVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }
}

