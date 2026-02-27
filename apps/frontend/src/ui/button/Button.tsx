import type { ComponentProps } from "react";
import { cn } from "../../util/cn";

// eslint-disable-next-line @typescript-eslint/no-empty-object-type
interface Props extends ComponentProps<"button"> {
  //
}
export function Button({ disabled, className, ...rest }: Props) {
  return (
    <button
      className={cn(
        "px-4 py-2 rounded-md bg-primary-500 text-black",
        disabled ? "opacity-50 cursor-not-allowed" : "",
        className,
      )}
      {...rest}
      disabled={disabled}
    />
  );
}
