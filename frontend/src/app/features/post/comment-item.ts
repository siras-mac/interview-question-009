import { ChangeDetectionStrategy, Component, input } from '@angular/core';

import { Comment } from '../../models/comment.model';

@Component({
  selector: 'app-comment-item',
  changeDetection: ChangeDetectionStrategy.OnPush,
  template: `
    <div class="comment">
      <span class="avatar">{{ comment().authorName.charAt(0) }}</span>
      <div class="comment-body">
        <span class="author">{{ comment().authorName }}</span>
        <p class="message">{{ comment().message }}</p>
      </div>
    </div>
  `,
  styles: `
    .comment {
      display: flex;
      gap: 12px;
      margin-bottom: 16px;
    }
    .avatar {
      flex-shrink: 0;
      width: 36px;
      height: 36px;
      border-radius: 50%;
      background: #3a66c5;
      color: #fff;
      font-weight: 700;
      display: flex;
      align-items: center;
      justify-content: center;
    }
    .author {
      font-weight: 700;
    }
    .message {
      margin: 2px 0 0;
    }
  `,
})
export class CommentItem {
  readonly comment = input.required<Comment>();
}
