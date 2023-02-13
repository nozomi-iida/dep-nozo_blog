import { useToast as useChakraToast, UseToastOptions } from "@chakra-ui/toast";

export const useCustomToast = () => {
  const toast = useChakraToast();

  const customToast = (options: UseToastOptions) => {
    toast({
      position: "top",
      ...options,
    });
  };

  return customToast;
};
