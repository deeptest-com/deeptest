package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.io.Serializable;

@Entity
@Table(name = "sys_nums")
public class SysNums implements Serializable {
    private static final long serialVersionUID = -9140817575372640340L;

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long key;

    public Long getKey() {
        return key;
    }

    public void setKey(Long key) {
        this.key = key;
    }
}
