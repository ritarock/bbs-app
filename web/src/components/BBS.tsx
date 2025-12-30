import { useState } from "react";
import { PostList } from "./PostList";
import { PostDetail } from "./PostDetail";

export function BBS() {
  const [selectedPostId, setSelectedPostId] = useState<number | null>(null);

  return (
    <div className="bbs">
      <h1 className="bbs-title">BBS</h1>
      <div className="bbs-container">
        <PostList
          selectedPostId={selectedPostId}
          onSelectPost={setSelectedPostId}
        />
        <PostDetail postId={selectedPostId} />
      </div>
    </div>
  );
}
