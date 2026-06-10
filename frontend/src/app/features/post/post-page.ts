import { ChangeDetectionStrategy, Component, OnInit, inject, signal } from '@angular/core';
import { DatePipe } from '@angular/common';
import { FormsModule } from '@angular/forms';

import { Comment } from '../../models/comment.model';
import { Post } from '../../models/post.model';
import { CommentService } from '../../services/comment.service';
import { CommentItem } from './comment-item';

const POST_ID = 1;

@Component({
  selector: 'app-post-page',
  changeDetection: ChangeDetectionStrategy.OnPush,
  imports: [DatePipe, FormsModule, CommentItem],
  templateUrl: './post-page.html',
  styleUrl: './post-page.css',
})
export class PostPage implements OnInit {
  private readonly commentService = inject(CommentService);

  protected readonly currentUser = 'Blend 285';
  protected readonly post = signal<Post | null>(null);
  protected readonly comments = signal<Comment[]>([]);
  protected draft = '';

  ngOnInit(): void {
    this.commentService.getPost(POST_ID).subscribe((post) => this.post.set(post));
    this.commentService.getComments(POST_ID).subscribe((comments) => this.comments.set(comments));
  }

  protected submit(): void {
    const message = this.draft.trim();
    if (!message) {
      return;
    }
    this.commentService.addComment(POST_ID, this.currentUser, message).subscribe((comment) => {
      this.comments.update((list) => [...list, comment]);
      this.draft = '';
    });
  }
}
