export const markdown2content = (str: string) => {
  const markdownRegex = /(\*\*|__|\*|_|~~|`|>|#+|\[|\]\(|\))/g;
  const urlRegex = /https?:\/\/[^\s]+/g;

  return str.replace(markdownRegex, "").replace(urlRegex, "");
};
