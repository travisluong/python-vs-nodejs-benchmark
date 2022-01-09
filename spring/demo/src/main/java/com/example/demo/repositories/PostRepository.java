package com.example.demo.repositories;

import com.example.demo.domain.Post;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource
public interface PostRepository extends PagingAndSortingRepository<Post, Long> {}
