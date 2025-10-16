# Kernels

Firecracker requires a kernel image to boot a microVM.

Currently, Firecracker supports uncompressed ELF kernel images on x86_64 while on aarch64 it supports PE formatted images.

For compatibility, it is recommended to use a kernel image that is supported by the Firecracker version you are using. For more information see the [Firecracker's Kernel Support Policy](https://github.com/firecracker-microvm/firecracker/blob/main/docs/kernel-policy.md)

With the default Fireactions installation, the kernel image is located in the `/var/lib/fireactions/kernels/<VERSION>/vmlinux` path. You can change the kernel image path by modifying the `kernel_image_path` parameter for each [Pool](./concepts.md#pool) in the Fireactions configuration file.

The latest Firecracker kernel image can be downloaded from the [Firecracker CI](https://s3.amazonaws.com/spec.ccfc.min/) S3 bucket. The following script can be used to download the latest kernel image for a specific version:

```bash
#!/bin/bash

latest=$(wget "http://spec.ccfc.min.s3.amazonaws.com/?prefix=firecracker-ci/v1.10/x86_64/vmlinux-5.10&list-type=2" -O - 2>/dev/null | grep "(?<=<Key>)(firecracker-ci/v1.10/x86_64/vmlinux-5\.10\.[0-9]{3})(?=</Key>)" -o -P)

wget "https://s3.amazonaws.com/spec.ccfc.min/${latest}"
```

We also provide our own Kernel images, customized for GitHub runners. The configuration can be found [here](https://github.com/hostinger/fireactions-images/tree/main/kernels). The kernel images can be downloaded from the following links:

* https://storage.googleapis.com/fireactions/kernels/amd64/5.10/vmlinux
* https://storage.googleapis.com/fireactions/kernels/arm64/5.10/vmlinux

You can also build your own custom kernel image. For more information see the [Firecracker documentation](https://github.com/firecracker-microvm/firecracker).
