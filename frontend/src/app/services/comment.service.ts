import { HttpClient } from '@angular/common/http';
import { Injectable, inject } from '@angular/core';
import { Observable } from 'rxjs';

import { Comment } from '../models/comment.model';
import { Post } from '../models/post.model';

@Injectable({ providedIn: 'root' })
export class CommentService {
  private readonly http = inject(HttpClient);
  private readonly baseUrl = '/api';

  getPost(id: number): Observable<Post> {
    return this.http.get<Post>(`${this.baseUrl}/posts/${id}`);
  }

  getComments(postId: number): Observable<Comment[]> {
    return this.http.get<Comment[]>(`${this.baseUrl}/posts/${postId}/comments`);
  }

  addComment(postId: number, authorName: string, message: string): Observable<Comment> {
    return this.http.post<Comment>(`${this.baseUrl}/posts/${postId}/comments`, {
      authorName,
      message,
    });
  }
}
