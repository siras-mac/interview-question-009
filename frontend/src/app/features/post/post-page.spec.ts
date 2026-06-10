import { TestBed } from '@angular/core/testing';
import { provideHttpClient } from '@angular/common/http';
import { HttpTestingController, provideHttpClientTesting } from '@angular/common/http/testing';

import { PostPage } from './post-page';

const mockPost = {
  id: 1,
  authorName: 'Change can',
  imageUrl: 'assets/post.png',
  postedAt: '2021-10-16T16:00:00+07:00',
};

const seedComment = {
  id: 1,
  postId: 1,
  authorName: 'Blend 285',
  message: 'have a good day',
  createdAt: '2021-10-16T16:05:00+07:00',
};

describe('PostPage', () => {
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PostPage],
      providers: [provideHttpClient(), provideHttpClientTesting()],
    }).compileComponents();
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => httpMock.verify());

  async function createPage() {
    const fixture = TestBed.createComponent(PostPage);
    fixture.detectChanges();

    httpMock.expectOne('/api/posts/1').flush(mockPost);
    httpMock.expectOne('/api/posts/1/comments').flush([seedComment]);
    await fixture.whenStable();
    fixture.detectChanges();
    return fixture;
  }

  it('renders the post and existing comments', async () => {
    const fixture = await createPage();
    const el = fixture.nativeElement as HTMLElement;

    expect(el.querySelector('.page-header')?.textContent).toContain('IT 08-1');
    expect(el.textContent).toContain('Change can');
    expect(el.textContent).toContain('have a good day');
  });

  it('adds a comment below the list when pressing ENTER', async () => {
    const fixture = await createPage();
    const el = fixture.nativeElement as HTMLElement;
    const input = el.querySelector<HTMLInputElement>('.comment-input')!;

    input.value = 'nice photo';
    input.dispatchEvent(new Event('input'));
    await fixture.whenStable();

    input.dispatchEvent(new KeyboardEvent('keydown', { key: 'Enter' }));

    const req = httpMock.expectOne('/api/posts/1/comments');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ authorName: 'Blend 285', message: 'nice photo' });
    req.flush({
      id: 2,
      postId: 1,
      authorName: 'Blend 285',
      message: 'nice photo',
      createdAt: '2026-06-10T12:00:00+07:00',
    });

    await fixture.whenStable();
    fixture.detectChanges();

    const comments = el.querySelectorAll('app-comment-item');
    expect(comments.length).toBe(2);
    expect(comments[1].textContent).toContain('nice photo');
    expect(input.value).toBe('');
  });

  it('does not send a request for a blank comment', async () => {
    const fixture = await createPage();
    const el = fixture.nativeElement as HTMLElement;
    const input = el.querySelector<HTMLInputElement>('.comment-input')!;

    input.value = '   ';
    input.dispatchEvent(new Event('input'));
    await fixture.whenStable();

    input.dispatchEvent(new KeyboardEvent('keydown', { key: 'Enter' }));

    httpMock.expectNone('/api/posts/1/comments');
  });
});
