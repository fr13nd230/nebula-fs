type Data<T> = {
  content: T | T[];
};

export type Base = {
  status: boolean;
  code: number;
  message: string;
};

export type Response<T> = Base & Data<T>;
