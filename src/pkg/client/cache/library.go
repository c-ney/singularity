// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cache

// Library returns the directory inside the cache.Dir() where library images are cached
func Library() string {
	return updateCacheSubdir(LibraryDir)
}
