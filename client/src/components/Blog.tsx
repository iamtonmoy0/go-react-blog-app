import axios from "axios";
import React, { useEffect, useState } from "react";

export default function Blog() {
  const [blogs, setBlogs] = useState([]);
  useEffect(() => {
    axios
      .get("http://localhost:3000/all-blog")
      .then((res) => {
        setBlogs(res.data.data);
        // console.log(res.data.data);
      })
      .catch((error) => {
        console.log(error);
      });
  }, []);

  return (
    <div className="px-5 pt-5 w-full grid grid-cols-3 gap-4">
      {blogs &&
        blogs.map((blog) => (
          <article
            key={blog.id}
            className="hover:animate-background rounded-xl bg-gradient-to-r from-green-300 via-blue-500 to-purple-600 p-0.5 shadow-xl transition hover:bg-[length:400%_400%] hover:shadow-sm hover:[animation-duration:_4s]"
          >
            <div className="rounded-[10px] bg-white p-4 !pt-20 sm:p-6">
              <time
                dateTime="2022-10-10"
                className="block text-xs text-gray-500"
              >
                ID:{blog.id}
              </time>

              <a href="#">
                <h3 className="mt-0.5 text-lg font-medium text-gray-900">
                  {blog.title}
                </h3>
              </a>
              <p>{blog.post}</p>
            </div>
          </article>
        ))}
    </div>
  );
}
