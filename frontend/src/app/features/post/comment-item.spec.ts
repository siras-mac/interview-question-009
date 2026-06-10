import { TestBed } from '@angular/core/testing';

import { CommentItem } from './comment-item';

describe('CommentItem', () => {
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CommentItem],
    }).compileComponents();
  });

  it('renders author name, message and avatar initial', async () => {
    const fixture = TestBed.createComponent(CommentItem);
    fixture.componentRef.setInput('comment', {
      id: 1,
      postId: 1,
      authorName: 'Blend 285',
      message: 'have a good day',
      createdAt: '2021-10-16T16:05:00+07:00',
    });
    await fixture.whenStable();

    const el = fixture.nativeElement as HTMLElement;
    expect(el.querySelector('.author')?.textContent).toBe('Blend 285');
    expect(el.querySelector('.message')?.textContent).toBe('have a good day');
    expect(el.querySelector('.avatar')?.textContent).toBe('B');
  });
});
