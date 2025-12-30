const TOKEN_KEY = "bbs_auth_token";

export const customFetch = async <T>({
  url,
  method,
  params,
  data,
  signal,
}: {
  url: string;
  method: "GET" | "POST" | "PUT" | "DELETE" | "PATCH";
  params?: Record<string, string>;
  data?: unknown;
  headers?: Record<string, string>;
  signal?: AbortSignal;
}): Promise<T> => {
  const searchParams = params
    ? `?${new URLSearchParams(params).toString()}`
    : "";
  const fullUrl = `${url}${searchParams}`;

  const token = localStorage.getItem(TOKEN_KEY);
  const headers: Record<string, string> = {
    "Content-Type": "application/json",
  };
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const response = await fetch(fullUrl, {
    method,
    headers,
    body: data ? JSON.stringify(data) : undefined,
    signal,
  });

  if (!response.ok) {
    const error = await response.json().catch(() => ({}));

    if (response.status === 401) {
      localStorage.removeItem(TOKEN_KEY);
      localStorage.removeItem("bbs_user");
    }

    throw new Error(error.message || `HTTP Error: ${response.status}`);
  }

  if (response.status === 204) {
    return undefined as T;
  }

  return response.json();
};
