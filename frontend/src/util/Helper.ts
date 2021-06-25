export const Helper = {
  traslit: (x: string) => {
    const a = '`qwertyuiop[]asdfghjkl;\'zxcvbnm,.QWERTYUIOP{}ASDFGHJKL:"ZXCVBNM<>';
    const b = 'ёйцукенгшщзхъфывапролджэячсмитьбюЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮ';
    const out = x.split('');
    for (let i = 0; i < x.length; i++) {
      for (let j = 0; j < a.length; j++) {
        if (x[i] === a[j]) {
          out[i] = b[j];
          break;
        }
      }
    }
    return out.join('');
  },
};
