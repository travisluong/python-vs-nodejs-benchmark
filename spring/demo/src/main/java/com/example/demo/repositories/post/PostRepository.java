package com.example.demo.repositories.post;

import com.example.demo.domain.post.Post;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource
public interface PostRepository extends PagingAndSortingRepository<Post, Long> {}
