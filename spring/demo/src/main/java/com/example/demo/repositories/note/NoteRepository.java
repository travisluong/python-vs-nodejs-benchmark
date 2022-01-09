package com.example.demo.repositories.note;

import com.example.demo.domain.note.Note;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource
public interface NoteRepository extends PagingAndSortingRepository<Note, Long> {
    
}
