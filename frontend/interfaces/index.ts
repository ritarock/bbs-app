export type Topic = {
  id: number;
  name: string;
  detail: string;
  created_at: Date;
  updated_at: Date;
};

export type Comment = {
  id: number;
  body: string;
  created_at: Date;
  updated_at: Date;
  topic_id: number;
};
