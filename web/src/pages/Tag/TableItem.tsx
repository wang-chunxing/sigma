/**
 * Copyright 2023 XImager
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */



import { useRef, useState } from "react";
import { useClickAway } from 'react-use';
import { useNavigate } from 'react-router-dom';
import relativeTime from 'dayjs/plugin/relativeTime';
import dayjs from 'dayjs';

dayjs.extend(relativeTime);

export default function TableItem({ id, namespace, repository, name, digest, size, created_at, updated_at }: { id: number, namespace: string | undefined, repository: string | null, name: string, digest: string, size: number, created_at: string, updated_at: string }) {
  const navigate = useNavigate();
  let [show, setShow] = useState(false);

  const ref = useRef<HTMLDivElement>() as React.MutableRefObject<HTMLDivElement>;;
  useClickAway(ref, () => {
    if (show) {
      setShow(!show);
    }
  });

  return (
    <tr className="cursor-pointer">
      <td className="px-6 py-4 max-w-0 w-full whitespace-nowrap text-sm font-medium text-gray-900">
        <div className="flex items-center space-x-3 lg:pl-2">
          <div className="cursor-pointer truncate hover:text-gray-600">
            {name}
          </div>
        </div>
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {digest}
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {size}
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {dayjs().to(dayjs(created_at))}
      </td>
      <td className="hidden md:table-cell px-6 py-4 whitespace-nowrap text-sm text-gray-500 text-right">
        {dayjs().to(dayjs(updated_at))}
      </td>
      <td className="pr-3 whitespace-nowrap">
        <button
          type="button"
          className=" w-1/2  rounded-md border border-transparent bg-white font-medium text-indigo-600 hover:text-indigo-500  mr-5"
        >
          Remove
        </button>
      </td>
    </tr>
  );
}