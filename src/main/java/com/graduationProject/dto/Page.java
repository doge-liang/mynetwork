package com.graduationProject.dto;

import lombok.Data;

/**
 * Page
 * <p>
 * 页面类
 * <p>
 * Created : 2021/5/17 15:45
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@Data
public class Page<T> {

    private T data;
    private Integer pageSize;
    private String bookmark;

    public Page(T data, String bookmark){
        this.data = data;
        this.bookmark = bookmark;
        this.pageSize = 40;
    }
}
