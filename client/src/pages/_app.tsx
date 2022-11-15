import type { AppProps } from "next/app";
import { ChakraProvider, ColorModeScript } from "@chakra-ui/react";
import chakraTheme from "../libs/chakra/theme";
import { NextPage } from "next";
import { ReactElement, ReactNode } from "react";
import Head from "next/head";

export type NextPageWithLayout<P = {}, IP = P> = NextPage<P, IP> & {
  getLayout?: (page: ReactElement) => ReactNode;
};

type AppPropsWithLayout = AppProps & {
  Component: NextPageWithLayout;
};

export const App = ({ Component, pageProps }: AppPropsWithLayout) => {
  const getLayout = Component.getLayout || ((page) => page);
  return (
    <div>
      <Head>
        <title>Nozo Blog</title>
        <meta
          http-equiv="Content-Security-Policy"
          content="upgrade-insecure-requests"
        />
      </Head>
      <ChakraProvider theme={chakraTheme}>
        <ColorModeScript initialColorMode={chakraTheme.config.theme} />
        {getLayout(<Component {...pageProps} />)}
      </ChakraProvider>
    </div>
  );
};

export default App;
