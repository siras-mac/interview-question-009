import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';

import { CommentService } from './comment.service';

describe('CommentService', () => {
  let service: CommentService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [provideHttpClient(), provideHttpClientTesting()],
    });
    service = TestBed.inject(CommentService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => httpMock.verify());

  it('gets a post by id', () => {
    const mockPost = {
      id: 1,
      authorName: 'Change can',
      imageUrl: 'assets/post.png',
      postedAt: '2021-10-16T16:00:00+07:00',
    };

    service.getPost(1).subscribe((post) => expect(post).toEqual(mockPost));

    const req = httpMock.expectOne('/api/posts/1');
    expect(req.request.method).toBe('GET');
    req.flush(mockPost);
  });

  it('gets comments of a post', () => {
    const mockComments = [
      {
        id: 1,
        postId: 1,
        authorName: 'Blend 285',
        message: 'have a good day',
        createdAt: '2021-10-16T16:05:00+07:00',
      },
    ];

    service.getComments(1).subscribe((comments) => expect(comments).toEqual(mockComments));

    const req = httpMock.expectOne('/api/posts/1/comments');
    expect(req.request.method).toBe('GET');
    req.flush(mockComments);
  });

  it('posts a new comment', () => {
    const mockComment = {
      id: 2,
      postId: 1,
      authorName: 'Blend 285',
      message: 'nice photo',
      createdAt: '2026-06-10T12:00:00+07:00',
    };

    service
      .addComment(1, 'Blend 285', 'nice photo')
      .subscribe((comment) => expect(comment).toEqual(mockComment));

    const req = httpMock.expectOne('/api/posts/1/comments');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ authorName: 'Blend 285', message: 'nice photo' });
    req.flush(mockComment);
  });
});
