const get = async <T>(path: string, token: string): Promise<T> => {
  return fetch(path, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token,
    }
  }).then<T>((res) => res.json())
}

const post = async <T>(path: string, data: T, token: string): Promise<T> => {
  const res = await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + token,
    },
    body: JSON.stringify(data)
  })

  return res.json()
}

const postNoToken = async <T>(path: string, data: T) => {
  const res = await fetch(path, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(data),
  });

  return res.json()
}

export { get, post, postNoToken }
